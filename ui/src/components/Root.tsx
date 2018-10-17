/// <reference path="../../node_modules/@types/reactstrap/index.d.ts" />
import * as React from 'react';

import 'bootstrap/dist/css/bootstrap.min.css';
import { Container, Row, Col, Button } from 'reactstrap';

import MemoList from './MemoList'
import MemoEdit from './MemoEdit'
import { Memo } from '../api'

interface RootProps {
  memos: Memo[];
  onRefresh(): void;
  onSave(content: string): void;
}

const Root: React.SFC<RootProps> = ({memos, onRefresh, onSave}) => {
  console.log("Root memos", memos)
  return (
    <Container>
      <Row>
        <Col md="6">
          <MemoList memos={memos}/>
          <Button color="primary" tag="button" onClick={onRefresh}>Refresh</Button>
        </Col>
        <Col md="6"><MemoEdit onSave={onSave}/></Col>
      </Row>
    </Container>
  )
};

export default Root
