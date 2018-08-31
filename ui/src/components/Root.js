// src/components/Root.js
import React from 'react';

import 'bootstrap/dist/css/bootstrap.min.css';
import { Container, Row, Col } from 'reactstrap';

import MemoList from './MemoList'
import MemoEdit from './MemoEdit'

import client from '../api/client'

const Root = () => {
  client().listMemos("/memos")
    .then( (resp) => console.log("SUCCESS response", resp.statusText) )
    .catch( (resp) => console.log("ERROR response", resp.statusText) );

  return (
    <Container>
      <Row>
        <Col md="6"><MemoList /></Col>
        <Col md="6"><MemoEdit /></Col>
      </Row>
    </Container>
  )
};

export default Root
