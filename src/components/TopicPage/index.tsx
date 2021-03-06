import React, { Component } from 'react'
import { graphql, createFragmentContainer, RelayProp } from 'react-relay'
import { isEmpty } from 'ramda'
import classNames from 'classnames'
import { Link as FoundLink } from 'found'

import Page from 'components/ui/Page'
import Subhead from 'components/ui/Subhead'
import Breadcrumbs from 'components/ui/Breadcrumbs'
import SidebarList from 'components/ui/SidebarList'
import Columns from 'components/ui/Columns'
import LeftColumn from 'components/ui/LeftColumn'
import RightColumn from 'components/ui/RightColumn'
import List from 'components/ui/List'
import Link from 'components/ui/Link'
import Topic from 'components/ui/Topic'
import { LocationType, NodeTypeOf, liftNodes } from 'components/types'
import { TopicPage_query_QueryResponse as Response } from '__generated__/TopicPage_query_Query.graphql'
import { TopicPage_topic as TopicType } from '__generated__/TopicPage_topic.graphql'
import AddForm from './AddForm'
import styles from './styles.module.css'

type ViewType = Response['view']
type LinkType = NodeTypeOf<TopicType['links']>
type ChildTopicType = NodeTypeOf<TopicType['childTopics']>
type ParentTopicType = NodeTypeOf<TopicType['parentTopics']>

type Props = {
  // eslint-disable-next-line react/no-unused-prop-types
  alerts: Object[],
  location: LocationType,
  orgLogin: string,
  relay: RelayProp,
  router: Object,
  topic: TopicType,
  view: ViewType,
}

type State = {}

class TopicPage extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {}
  }

  static getDerivedStateFromProps = (nextProps: Props) => {
    const shouldAppend = window.flashMessages && nextProps.alerts && nextProps.alerts.length > 0
    if (shouldAppend && window.flashMessages?.addMessage)
      // eslint-disable-next-line nonblock-statement-body-position
      nextProps.alerts.forEach(window.flashMessages.addMessage)

    return {}
  }

  get links() {
    return liftNodes<LinkType>(this.props.topic.links)
  }

  get topics() {
    return liftNodes<ChildTopicType>(this.props.topic.childTopics)
  }

  get synonyms() {
    return this.props.topic.synonyms
  }

  get isGuest(): boolean {
    return this.props.view.viewer.isGuest
  }

  get repoName(): string {
    const { currentRepository } = this.props.view
    return currentRepository ? currentRepository.displayName : 'No name'
  }

  get recentActivityLocation(): LocationType {
    return {
      pathname: `${this.props.topic.resourcePath}/recent`,
      query: {},
      search: '',
      state: {
        orgLogin: this.props.orgLogin,
        repoName: this.repoName,
        itemTitle: this.props.topic.displayName,
      },
    }
  }

  get linksToReviewLocation(): LocationType {
    return {
      pathname: `${this.props.topic.resourcePath}/review`,
      query: {},
      search: '',
      state: {
        orgLogin: this.props.orgLogin,
        repoName: this.repoName,
        itemTitle: this.props.topic.displayName,
      },
    }
  }

  renderLink = (link: LinkType | null) => (
    link && (
    <Link
      key={link.id}
      link={link}
      orgLogin={this.props.orgLogin}
      view={this.props.view}
      viewer={this.props.view.viewer}
    />
    )
  )

  renderTopic = (topic: ChildTopicType | null) => (
    topic && (
    <Topic
      key={topic.id}
      orgLogin={this.props.orgLogin}
      topic={topic}
      view={this.props.view}
    />
    )
  )

  renderAddForm = () => (
    <AddForm
      topic={this.props.topic}
      viewer={this.props.view.viewer}
    />
  )

  renderHeadingDetail = () => {
    const { synonyms } = this
    const { length } = synonyms

    if (length < 2) return null

    return (
      <div className={classNames(styles.synonyms, 'h6')}>
        {synonyms.slice(1, length).map(({ name }) => (
          <span key={name} className={styles.synonym}>{name}</span>
        ))}
      </div>
    )
  }

  renderNotification = () => (
    <div className="Box p-3 mt-3">
      You must be
      {' '}
      <a href="/login">signed in</a>
      {' '}
      to add and move topics and links.
    </div>
  )

  renderTopicViews = () => (
    <div className="Box Box--condensed mb-3">
      <div className="Box-header">
        <span className="Box-title">This topic</span>
      </div>
      <ul>
        <li className="Box-row">
          <FoundLink to={this.recentActivityLocation} className="Box-row-link">
            Recent activity
          </FoundLink>
        </li>
        {!this.isGuest && (
          <li className="Box-row">
            <FoundLink to={this.linksToReviewLocation} className="Box-row-link">
              Links to review
            </FoundLink>
          </li>
        )}
      </ul>
    </div>
  )

  render = () => {
    const { location, topic, view } = this.props

    if (!topic) {
      return (
        <div>
          Topic not found:
          {location.pathname}
        </div>
      )
    }

    const { displayName, parentTopics } = topic
    const { topics, links, repoName } = this
    const { currentRepository } = view

    return (
      <Page>
        <Breadcrumbs
          orgLogin={this.props.orgLogin}
          repository={currentRepository}
        />
        <Subhead
          heading={displayName}
          renderHeadingDetail={this.renderHeadingDetail}
        />
        <Columns>
          <RightColumn>
            <SidebarList
              items={liftNodes<ParentTopicType>(parentTopics)}
              orgLogin={this.props.orgLogin}
              placeholder="There are no parent topics for this topic."
              repoName={repoName}
              title="Parent topics"
            />
            { this.renderTopicViews() }
            { this.isGuest
              ? this.renderNotification()
              : this.renderAddForm()}
          </RightColumn>
          <LeftColumn>
            <List
              placeholder="There are no items in this list."
              hasItems={!isEmpty(topics) || !isEmpty(links)}
            >
              { topics.map(this.renderTopic) }
              { links.map(this.renderLink) }
            </List>
          </LeftColumn>
        </Columns>
      </Page>
    )
  }
}

export const UnwrappedTopicPage = TopicPage

export const query = graphql`
query TopicPage_query_Query(
  $viewerId: ID!,
  $orgLogin: String!,
  $repoName: String,
  $repoIds: [ID!],
  $topicId: ID!,
  $searchString: String,
) {
  alerts {
    id
    text
    type
  }

  view(
    viewerId: $viewerId,
    currentOrganizationLogin: $orgLogin,
    currentRepositoryName: $repoName,
    repositoryIds: $repoIds,
  ) {
    viewer {
      id
      isGuest
      ...AddForm_viewer
      ...Link_viewer
    }

    currentRepository {
      displayName
      ...Breadcrumbs_repository
    }

    ...Link_view
    ...Topic_view

    topic(id: $topicId) {
      ...TopicPage_topic @arguments(searchString: $searchString)
    }
  }
}`

export default createFragmentContainer(TopicPage, {
  topic: graphql`
    fragment TopicPage_topic on Topic @argumentDefinitions(
      searchString: {type: "String", defaultValue: ""},
    ) {
      displayName: name
      id
      resourcePath

      synonyms {
        name
      }

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

      ...AddForm_topic
    }
  `,
})
