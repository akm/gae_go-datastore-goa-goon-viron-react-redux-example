// src/components/Root.js
import React from 'react';

import 'bootstrap/dist/css/bootstrap.min.css';
import { Container, Row, Col, Button } from 'reactstrap';

import MemoList from './MemoList'
import MemoEdit from './MemoEdit'


const Root = ({memos, onRefresh}) => {
  console.log("Root memos", memos)
  return (
    <Container>
      <Row>
        <Col md="6">
          <MemoList memos={memos}/>
          <Button color="primary" tag="button" onClick={onRefresh}>Refresh</Button>
        </Col>
        <Col md="6"><MemoEdit /></Col>
      </Row>
    </Container>
  )
};

export default Root
