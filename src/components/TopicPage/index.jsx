// @flow
import React, { Component } from 'react'
import { graphql, createFragmentContainer } from 'react-relay'
import { isEmpty } from 'ramda'

import { liftNodes } from 'utils'
import type { LinkType, Relay, TopicType, UserType, ViewType } from 'components/types'
import Subhead from 'components/ui/Subhead'
import SidebarList from 'components/ui/SidebarList'
import List from 'components/ui/List'
import Link from 'components/ui/Link'
import Topic from 'components/ui/Topic'
import Breadcrumbs from 'components/ui/Breadcrumbs'
import AddForm from './AddForm'

type Props = {
  location: Object,
  orgLogin: string,
  relay: Relay,
  router: Object,
  topic: TopicType,
  view: ViewType,
  viewer: UserType,
}

class TopicPage extends Component<Props> {
  get links(): LinkType[] {
    return liftNodes(this.props.topic.links)
  }

  get topics(): TopicType[] {
    return liftNodes(this.props.topic.childTopics)
  }

  render = () => {
    const { location, topic } = this.props

    if (!topic)
      return <div>Topic not found: {location.pathname}</div>

    const { name, parentTopics, resourcePath } = topic
    const { topics, links } = this

    return (
      <div>
        <Breadcrumbs
          orgLogin={this.props.orgLogin}
          repository={this.props.view.currentRepository}
        />
        <Subhead
          heading={name}
          headingLink={resourcePath}
          location={this.props.location}
          router={this.props.router}
          view={this.props.view}
        />
        <div className="two-thirds column pl-0">
          <List
            placeholder="There are no items in this list."
            hasItems={!isEmpty(topics) || !isEmpty(links)}
          >
            { topics.map(childTopic => (
              <Topic
                key={childTopic.id}
                orgLogin={this.props.orgLogin}
                relay={this.props.relay}
                topic={childTopic}
                view={this.props.view}
              />
            )) }

            { links.map(link => (
              <Link
                key={link.id}
                link={link}
                orgLogin={this.props.orgLogin}
                relay={this.props.relay}
                view={this.props.view}
              />
            )) }
          </List>
        </div>
        <div className="one-third column pr-0">
          <SidebarList
            title="Parent topics"
            items={liftNodes(parentTopics)}
          />
          <AddForm
            relay={this.props.relay}
            topic={topic}
            viewer={this.props.viewer}
          />
        </div>
      </div>
    )
  }
}

export const query = graphql`
query TopicPage_query_Query(
  $orgLogin: String!,
  $repoName: String,
  $repoIds: [ID!],
  $topicId: ID!,
  $searchString: String,
) {
  viewer {
    id
    ...AddForm_viewer
  }

  view(
    currentOrganizationLogin: $orgLogin,
    currentRepositoryName: $repoName,
    repositoryIds: $repoIds,
  ) {
    currentRepository {
      ...Breadcrumbs_repository
    }

    ...Link_view
    ...Topic_view
    ...Subhead_view

    topic(id: $topicId) {
      ...TopicPage_topic @arguments(searchString: $searchString)
    }
  }
}`

export default createFragmentContainer(TopicPage, graphql`
  fragment TopicPage_topic on Topic @argumentDefinitions(
    searchString: {type: "String", defaultValue: ""},
  ) {
    name
    resourcePath
    ...AddForm_topic

    parentTopics(first: 100) {
      edges {
        node {
          display: name
          resourcePath
        }
      }
    }

    childTopics(first: 1000, searchString: $searchString) @connection(key: "Topic_childTopics") {
      edges {
        node {
          id
          ...Topic_topic
        }
      }
    }

    links(first: 1000, searchString: $searchString)  @connection(key: "Topic_links") {
      edges {
        node {
          id
          ...Link_link
        }
      }
    }
  }
`)
