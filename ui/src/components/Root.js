// src/components/Root.js
import React from 'react';

import client from '../api/client'

const Root = () => {
  client().listMemos("/memos")
    .then( (resp) => console.log("SUCCESS response", resp.statusText) )
    .catch( (resp) => console.log("ERROR response", resp.statusText) );

  return (
    <div>
        Hello World
    </div>
  )
};

export default Root
