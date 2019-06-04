// @flow
import React from 'react'
import type { Node } from 'react'
import { graphql } from 'react-relay'

import type { AlertType } from 'components/types'
import Header from './Header'
import FlashMessages from '../FlashMessages'

type Props = {
  alerts: ?AlertType[],
  children?: Node,
  defaultOrganization: Object,
  viewer: Object,
}

export const query = graphql`
query LayoutQuery {
  alerts {
    id
    text
    type
  }

  defaultOrganization {
    defaultRepository {
      rootTopic {
        resourcePath
      }
    }
  }

  viewer {
    ...Header_viewer
  }
}`

/* eslint jsx-a11y/anchor-is-valid: 0 */

const Layout = ({ alerts, children, defaultOrganization, viewer }: Props) => (
  <div className="layout">
    <Header
      className="clearfix mb-3 d-flex px-3 px-md-6 px-lg-4 py-2 box-shadow"
      viewer={viewer}
      defaultOrganization={defaultOrganization}
    />
    <div className="container-lg clearfix">
      <FlashMessages initialAlerts={alerts} />
      { children }
    </div>
    <footer>
      <div className="container-lg px-3 px-md-6 px-lg-0 my-6 pt-2 border-top">
        <p className="mb-2">
          Available for use under the MIT
          {' '}
          <a href="https://github.com/emwalker/digraph/blob/master/LICENSE.md">license</a>
.
          © Eric Walker.
        </p>
      </div>
    </footer>
  </div>
)

Layout.defaultProps = {
  children: null,
}

export default Layout
