// @flow
import React, { Component } from 'react'
import { createFragmentContainer, graphql } from 'react-relay'
import { pathOr } from 'ramda'

import SearchBox from 'components/ui/SearchBox'

const resourcePath = pathOr('/', ['defaultRepository', 'rootTopic', 'resourcePath'])

type Props = {
  heading: string,
  location: {
    pathname: string,
    query: Object,
    search: string,
  },
  router: {
    push: Function,
  },
  viewer: {
    defaultRepository: {
      rootTopic: {
        resourcePath: string,
      },
    },
  },
}

class Subhead extends Component<Props> {
  onSearch = (query) => {
    if (query === '') {
      this.props.router.push({ pathname: this.pathname })
      return
    }

    this.props.router.push({ pathname: this.pathname, query: { q: query } })
  }

  get pathname(): string {
    return resourcePath(this.props.viewer)
  }

  get searchString(): string {
    return this.props.location.search
      ? this.props.location.query.q
      : ''
  }

  render = () => (
    <div className="Subhead">
      <div className="Subhead-heading">{this.props.heading}</div>
      <SearchBox onEnter={this.onSearch} value={this.searchString} />
    </div>
  )
}

export default createFragmentContainer(Subhead, graphql`
  fragment Subhead_viewer on User {
    defaultRepository {
      rootTopic {
        resourcePath
      }
    }
  }
`)