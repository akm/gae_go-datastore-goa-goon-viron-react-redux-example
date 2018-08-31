import React from 'react';
import { Form, FormGroup, Label, Input } from 'reactstrap';

const MemoEdit = () => (
  <Form>
    <FormGroup>
       <Label for="exampleText">Text Area</Label>
       <Input type="textarea" name="text" id="exampleText" />
    </FormGroup>
  </Form>
)

export default MemoEdit
