import * as React from 'react';
import { Form, FormGroup, Label, Input, Button } from 'reactstrap';

interface MemoEditProps {
  onSave(content: string): void,
}
interface MemoEditState {
  content: string
}

class MemoEdit extends React.Component<MemoEditProps, MemoEditState> {
  constructor(props) {
    super(props);
    this.state = {content: ''};
  }

  handleChange = (event: React.FormEvent<HTMLInputElement>): void => {
    this.setState({content: event.currentTarget.value});
  }

  handleSubmit = (event:  React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
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
