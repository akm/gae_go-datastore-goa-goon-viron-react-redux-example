import React from 'react';
import { ListGroup, ListGroupItem } from 'reactstrap';

const MemoList = ({memos}) => {
  console.log("MemoList memos", memos)
  return (
    <ListGroup>
      { memos.map((memo) => {
        return <ListGroupItem>{memo}</ListGroupItem>;
      }) }
    </ListGroup>
  )
}

export default MemoList
