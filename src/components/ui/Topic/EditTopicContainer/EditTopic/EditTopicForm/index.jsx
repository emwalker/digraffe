// @flow
import React, { Component } from 'react'
import { createRefetchContainer, graphql } from 'react-relay'

import type { Option, TopicType } from 'components/types'
import deleteTopicMutation from 'mutations/deleteTopicMutation'
import updateTopicMutation from 'mutations/updateTopicMutation'
import updateTopicTopicsMutation from 'mutations/updateTopicParentTopicsMutation'
import EditTopicList, { makeOptions } from 'components/ui/EditTopicList'
import DeleteButton from 'components/ui/DeleteButton'
import Synonyms from './Synonyms'

type Props = {
  isOpen: boolean,
  orgLogin: string,
  relay: {
    environment: Object,
    refetch: Function,
  },
  toggleForm: Function,
  // $FlowFixMe
  topic: TopicType,
}

type State = {
  description: ?string,
  displayName: string,
}

class EditTopicForm extends Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      displayName: props.topic.displayName,
      description: props.topic.description,
    }
  }

  onSave = () => {
    updateTopicMutation(
      this.props.relay.environment,
      [],
      {
        topicIds: this.addTopicIds,
        description: this.state.description || '',
        id: this.props.topic.id,
        displayName: this.state.displayName,
      },
    )
    this.props.toggleForm()
  }

  onDelete = () => {
    deleteTopicMutation(
      this.props.relay.environment,
      [{
        type: 'NODE_DELETE',
        deletedIDFieldName: 'deletedTopicId',
      }],
      {
        topicId: this.props.topic.id,
      },
    )
  }

  // eslint-disable-next-line class-methods-use-this
  get addTopicIds(): string[] {
    return []
  }

  get topicId(): string {
    return this.props.topic.id
  }

  updateParentTopics = (parentTopicIds: string[]) => {
    updateTopicTopicsMutation(
      this.props.relay.environment,
      [],
      {
        topicId: this.props.topic.id,
        parentTopicIds,
      },
    )
  }

  updateDescription = (event: Object) => {
    this.setState({ description: event.currentTarget.value })
  }

  updateName = (event: Object) => {
    this.setState({ displayName: event.currentTarget.value })
  }

  loadOptions = (searchString: string): Promise<Option[]> => {
    if (!this.props.relay) return new Promise(() => [])

    return new Promise((resolve) => {
      const variables = {
        orgLogin: this.props.orgLogin,
        count: 60,
        searchString,
      }

      this.props.relay.refetch(variables, null, () => {
        const options = makeOptions(this.props.topic.availableTopics)
        resolve(options)
      })
    })
  }

  render = () => {
    if (!this.props.isOpen) return null

    return (
      <div className="my-4">
        <Synonyms relay={this.props.relay} topic={this.props.topic} />

        <EditTopicList
          loadOptions={this.loadOptions}
          selectedTopics={makeOptions(this.props.topic.selectedTopics)}
          updateTopics={this.updateParentTopics}
        />

        <dl className="form-group">
          <DeleteButton
            onDelete={this.onDelete}
          />
          <button
            className="btn-link float-right"
            onClick={this.props.toggleForm}
            type="button"
          >
            Close
          </button>
        </dl>
      </div>
    )
  }
}

export default createRefetchContainer(EditTopicForm, {
  viewer: graphql`
    fragment EditTopicForm_viewer on User {
      defaultRepository {
        id
      }
    }
  `,
  topic: graphql`
    fragment EditTopicForm_topic on Topic @argumentDefinitions(
      searchString: {type: "String", defaultValue: null},
      count: {type: "Int!", defaultValue: 10}
    ) {
      description
      id
      displayName

      selectedTopics: parentTopics(first: 100) {
        edges {
          node {
            value: id
            label: displayName
          }
        }
      }

      availableTopics: availableParentTopics(first: $count, searchString: $searchString) {
        edges {
          node {
            value: id
            label: displayName
          }
        }
      }

      ...Synonyms_topic
    }
  `,
},
graphql`
  query EditTopicFormRefetchQuery(
    $viewerId: ID!,
    $orgLogin: String!,
    $repoName: String,
    $repoIds: [ID!],
    $topicId: ID!,
    $count: Int!,
    $searchString: String,
  ) {
    view(
      viewerId: $viewerId,
      currentOrganizationLogin: $orgLogin,
      currentRepositoryName: $repoName,
      repositoryIds: $repoIds,
    ) {
      topic(id: $topicId) {
        ...EditTopicForm_topic @arguments(count: $count, searchString: $searchString)
      }
    }
  }
`)
