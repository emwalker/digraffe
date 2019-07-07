// @flow
import React, { Component } from 'react'
import { createFragmentContainer, graphql } from 'react-relay'

import { type CollectionNode } from 'components/types'
import { type ReviewPage_view as View } from './__generated__/ReviewPage_view.graphql'
import Container from './Container'
import Review from './Review'

type Link = CollectionNode<$PropertyType<View, 'links'>>

type Props = {
  view: View,
}

const Placeholder = () => (
  <Container totalCount={0}>
    <div className="blankslate">
      <p>Searching the computers for links to review ...</p>
    </div>
  </Container>
)

class ReviewPage extends Component<Props> {
  get links(): $ReadOnlyArray<?Link> {
    const { view: { links } } = this.props
    const edges = links ? links.edges : null
    if (!edges) return []
    return edges.map(edge => edge && edge.node)
  }

  get totalCount(): number {
    const { view: { links } } = this.props
    return links ? links.totalCount : 0
  }

  renderReview = (link: ?Link) => link && <Review key={link.id} link={link} />

  renderNoLinks = () => (
    <div className="blankslate">
      <p>There are no links to review.</p>
    </div>
  )

  render = () => {
    const { links } = this

    return (
      <Container totalCount={this.totalCount}>
        { links.length > 0
          ? links.map(this.renderReview)
          : this.renderNoLinks()
        }
      </Container>
    )
  }
}

export const query = graphql`
query ReviewPage_query_Query(
  $viewerId: ID!,
  $sessionId: ID!,
  $orgLogin: String!,
  $repoName: String,
  $repoIds: [ID!],
) {
  alerts {
    id
    text
    type
  }

  view(
    viewerId: $viewerId,
    sessionId: $sessionId,
    currentOrganizationLogin: $orgLogin,
    currentRepositoryName: $repoName,
    repositoryIds: $repoIds,
  ) {
    ...ReviewPage_view
  }
}`

const Wrapper = createFragmentContainer(ReviewPage, {
  view: graphql`
    fragment ReviewPage_view on View {
      links(first: 1000, reviewed: false) {
        totalCount

        edges {
          node {
            id
            ...Review_link
          }
        }
      }
    }
  `,
})

export default ({ props }: { props: Props }) => (
  // eslint-disable-next-line react/prop-types
  props && props.view
    ? <Wrapper {...props} />
    : <Placeholder />
)
