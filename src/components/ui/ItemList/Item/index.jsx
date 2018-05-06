// @flow
import React, { Component } from 'react'
import { pathOr } from 'ramda'
import classNames from 'classnames'

/* eslint no-underscore-dangle: 0 */

const edges = pathOr([], ['edges'])

type BadgeProps = {
  node: {
    name: string,
    resourcePath: string,
  }
}

const renderTopic = ({ node: { name, resourcePath } }: BadgeProps) => (
  <a className="Box-row-link" href={resourcePath} key={resourcePath}>
    <span
      className="Label Label--outline Label--outline-gray"
    >
      {name}
    </span>
  </a>
)

type Props = {
  __typename: string,
  display: string,
  resourcePath: string,
  topics: Array[],
}

class Item extends Component<Props> {
  get className(): string {
    return classNames(
      'Item-row',
      'Box-row',
      { 'Box-row--topic': this.props.__typename === 'Topic' },
    )
  }

  render() {
    return (
      <li
        className={this.className}
        key={this.props.resourcePath}
      >
        <div>
          <a className="Box-row-link" href={this.props.resourcePath}>
            { this.props.display || this.props.resourcePath }
          </a>
        </div>
        <div
          className="mt-1 link-url branch-name css-truncate css-truncate-target"
        >
          {this.props.resourcePath}
        </div>
        <div>
          { edges(this.props.topics).map(renderTopic) }
        </div>
      </li>
    )
  }
}

export default Item
