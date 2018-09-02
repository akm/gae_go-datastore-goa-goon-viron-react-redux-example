import React from 'react';
import { Form, FormGroup, Label, Input, Button } from 'reactstrap';

class MemoEdit extends React.Component {
  constructor(props) {
    super(props);
    this.state = {content: ''};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    console.log('handleChange', event);
    this.setState({content: event.target.value});
  }

  handleSubmit(event) {
    this.props.onSave(this.state.content);
  }

  render() {
    return (
      <Form onSubmit={this.handleSubmit}>
        <FormGroup>
          <Label for="exampleText">Text Area</Label>
          <Input type="textarea" name="text"
            onChange={this.handleChange} defaultValue={this.state.content}/>
        </FormGroup>
        <Button>Save</Button>
      </Form>
    )
  }
}

export default MemoEdit
