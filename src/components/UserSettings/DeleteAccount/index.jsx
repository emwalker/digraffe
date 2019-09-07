// @flow
import React, { useState, useCallback } from 'react'
import { createFragmentContainer, graphql } from 'react-relay'
import { ButtonDanger } from '@primer/components'

import type { Relay } from 'components/types'
import deleteAccountMutation from 'mutations/deleteAccountMutation'
import type { DeleteAccount_view as View } from './__generated__/DeleteAccount_view.graphql'

declare var confirm: Function

type Props = {
  relay: Relay,
  view: View,
}

const DeleteAccount = ({ relay, view }: Props) => {
  const [mutationInFlight, setMutationInFlight] = useState(false)
  const { viewer } = view

  const onClick = useCallback(async () => {
    // eslint-disable-next-line no-restricted-globals
    if (!confirm('Are you sure you want to delete your account?')) return

    setMutationInFlight(true)
    await deleteAccountMutation(relay.environment, [], { userId: viewer.id })

    setTimeout(
      () => {
        document.location.replace('/logout')
      },
      5000,
    )
  }, [mutationInFlight, relay, viewer])

  return (
    <>
      <div className="Subhead">
        <div className="Subhead-heading Subhead-heading--danger">Delete account</div>
      </div>
      <p>
        Your user information and private repo will be permanently removed when you delete your
        account. Links and topics that you have added to the public collection will still be there,
        but your account will no longer be associated with them.
      </p>
      <ButtonDanger disabled={mutationInFlight} onClick={onClick}>
        Delete your account
      </ButtonDanger>
      <p className="mt-5">
        To revoke GitHub auth permission for this account, go to the
        {' '}
        <a href="https://github.com/settings/applications">
          Authorized OAuth Apps
        </a>
        {' '}
        menu and look for &quot;Digraph&quot;.
      </p>
      <p>
        If you do not revoke permission for this app, it will seem as though your account was not
        deleted the next time you attempt to log into Digraph.  In that event, your account will be
        silently recreated using the information that you authorized Digraph to use when logging in
        the first time.
      </p>
    </>
  )
}

export default createFragmentContainer(DeleteAccount, {
  view: graphql`
    fragment DeleteAccount_view on View {
      viewer {
        id
      }
    }
  `,
})
