import queryMiddleware from 'farce/lib/queryMiddleware'
import createRender from 'found/lib/createRender'
import makeRouteConfig from 'found/lib/makeRouteConfig'
import Route from 'found/lib/Route'
import { Resolver } from 'found-relay'
import React from 'react'
import { graphql } from 'react-relay'
import { Environment, Network, RecordSource, Store } from 'relay-runtime'

import Homepage, { query as homepageQuery } from './components/Homepage'
import TopicPage, { query as topicPageQuery } from './components/TopicPage'
import TopicSearchPage, { query as topicSearchPageQuery } from './components/TopicSearchPage'
import Layout from './components/Layout'
import withErrorBoundary from './components/withErrorBoundary'

export const historyMiddlewares = [queryMiddleware]

export function createResolver(fetcher) {
  const environment = new Environment({
    network: Network.create((...args) => fetcher.fetch(...args)),
    store: new Store(new RecordSource()),
  })
  return new Resolver(environment)
}

const renderTopicPage = ({ props, error }: any) => {
  if (error)
    return <div>There was a problem.</div>

  if (!props)
    return null

  if (!props.view)
    return <div>You must log in and select an organization first.</div>

  const { location, params, view } = props

  if (location.query.q) {
    return (
      <TopicSearchPage
        orgLogin={params.orgLogin}
        repoName={params.repoName}
        topic={view.topic}
        location={location}
        {...props}
      />
    )
  }

  return (
    <TopicPage
      location={location}
      orgLogin={params.orgLogin}
      repoName={params.repoName}
      topic={view.topic}
      {...props}
    />
  )
}

/* eslint function-paren-newline: 0 */
export const routeConfig = makeRouteConfig(
  <Route
    Component={Layout}
    path="/"
    query={
      graphql`
      query router_Query {
        viewer {
          name
          avatarUrl

          defaultRepository {
            rootTopic {
              resourcePath
            }
          }
        }
      }`
    }
    prepareVariables={(params, { location }) => {
      const { q } = location.query
      return {
        ...params,
        repoIds: [],
        searchString: q,
      }
    }}
  >
    <Route
      Component={withErrorBoundary(Homepage)}
      query={homepageQuery}
    />
    <Route path=":orgLogin">
      <Route path="topics">
        <Route
          path=":topicId"
          render={renderTopicPage}
          getQuery={({ location }) => (
            location.query.q
              ? topicSearchPageQuery
              : topicPageQuery
          )}
        />
      </Route>
    </Route>
  </Route>,
)

export const render = createRender({})
