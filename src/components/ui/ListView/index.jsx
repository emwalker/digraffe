// @flow
import React from 'react'
import type { Node } from 'react'
import { ListGroup } from 'reactstrap'

import Item from './Item'

type Props = {
  children: Node,
  items: Array<{
    id: string,
    display: string,
    resourcePath: string,
  }>,
  title: string,
}

export default ({ children, items, title }: Props) => (
  <div className="listview">
    <h1>{title}</h1>
    <div className="row">
      <div className="col">
        <ListGroup>
          {items.map(({ resourcePath, ...props }) =>
            <Item key={resourcePath} {...props} />)}
        </ListGroup>
      </div>
      <div className="col-5">
        { children }
      </div>
    </div>
  </div>
)