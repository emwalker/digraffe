// @flow
import React, { Component, Fragment } from 'react'
import { createFragmentContainer, graphql } from 'react-relay'
import { Link } from 'found'
import classNames from 'classnames'

import type { UserType } from 'components/types'
import { toEverything } from 'components/navigation'
import DigraphLogo from 'components/ui/icons/DigraphLogo'
import UserNav from './UserNav'
import SignIn from './SignIn'
import { header, primaryLogo, userNav } from './styles.module.css'

type Props = {
  className?: ?string,
  viewer: UserType,
}

class DesktopNav extends Component<Props> {
  static defaultProps = {
    className: '',
  }

  get className(): string {
    return classNames('Header', this.props.className)
  }

  get isGuest(): boolean {
    const { viewer } = this.props

    return viewer ? viewer.isGuest : true
  }

  renderGuestUserNav = () => (
    <Fragment>
      <SignIn />
    </Fragment>
  )

  renderUserNav = (viewer: UserType) => <UserNav viewer={viewer} />

  render = () => {
    const { viewer } = this.props

    return (
      <header className={classNames(header, 'd-flex')}>
        <nav className="flex-self-center d-inline-block">
          <h1 className="h3 text-normal">
            <Link
              to="/"
              className={classNames(primaryLogo, 'text-gray-dark n-link no-underline d-flex')}
            >
              <div className="mr-2 d-inline-block">
                <DigraphLogo height="28px" width="28px" />
              </div>

              Digraph
            </Link>
          </h1>
        </nav>
        <nav className={classNames(userNav, 'flex-self-center')}>
          <a
            className="text-gray-dark px-2"
            href="https://blog.digraph.app"
          >
            Blog
          </a>
          <Link
            className="text-gray-dark px-2"
            id="recent-activity"
            to="/recent"
          >
            Recent
          </Link>
          <Link
            className="text-gray-dark px-2"
            to={toEverything}
          >
            Everything
          </Link>
          {this.isGuest
            ? this.renderGuestUserNav()
            : this.renderUserNav(viewer)
          }
        </nav>
      </header>
    )
  }
}

export const UnwrappedDesktopNav = DesktopNav

export default createFragmentContainer(DesktopNav, {
  viewer: graphql`
    fragment DesktopNav_viewer on User {
      isGuest
      ...UserNav_viewer
    }
  `,
})
