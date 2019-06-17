// @flow
import React, { Component, Fragment } from 'react'
import { Link } from 'found'

import { toEverything } from 'components/navigation'
import type { UserType } from 'components/types'

type Props = {
  viewer: UserType,
}

class Menu extends Component<Props> {
  renderSignIn = () => (
    <Link
      className="menu-item p-3 text-gray-dark"
      to="/login"
    >
      Sign in
    </Link>
  )

  renderUserNav = () => (
    <Fragment>
      <Link className="menu-item text-gray-dark p-3" to="/review">
        Review
      </Link>
      <a className="menu-item text-gray-dark p-3" href="/logout/github">Sign out</a>
    </Fragment>
  )

  render = () => (
    <nav className="menu" aria-label="Person settings">
      <a
        className="menu-item text-gray-dark p-3"
        href="https://blog.digraph.app"
      >
        Blog
      </a>
      <Link
        className="menu-item text-gray-dark p-3"
        id="recent-activity"
        to="/recent"
      >
        Recent
      </Link>
      <Link
        className="menu-item text-gray-dark p-3"
        to={toEverything}
      >
        Everything
      </Link>

      { this.props.viewer.isGuest
        ? this.renderSignIn()
        : this.renderUserNav()
      }
    </nav>
  )
}

export default Menu