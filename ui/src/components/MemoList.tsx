/// <reference path="../../node_modules/@types/reactstrap/index.d.ts" />
import * as React from 'react';
import { ListGroup, ListGroupItem } from 'reactstrap';
import { Memo } from '../states/Memo'

interface MemoListProps {
  memos: Memo[]
}

const MemoList: React.SFC<MemoListProps> = ({memos}) => {
  console.log("MemoList memos", memos)
  return (
    <ListGroup>
      { memos.map((memo, idx) => {
        return <ListGroupItem key={idx}>{memo.content}</ListGroupItem>;
      }) }
    </ListGroup>
  )
}

export default MemoList
