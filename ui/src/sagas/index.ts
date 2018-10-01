import { delay } from 'redux-saga'
import { call, takeEvery, put } from 'redux-saga/effects'
import { AxiosResponse, AxiosError } from 'axios';

import client from '../api/client'
import { IndexActionType, SaveAsyncAction, addMemoAction, refreshPostAction } from '../actions/index'

export default function* rootSaga() {
  yield takeEvery(IndexActionType.REFRESH_ASYNC, refreshAsync)
  yield takeEvery(IndexActionType.SAVE_ASYNC, saveAsync)
}

export function* refreshAsync() {
  const memos = yield call(fetchMemos); // TODO handle error
  yield put(refreshPostAction(memos))
}

function fetchMemos() {
  return client().listMemos("/memos")
    .then((resp: AxiosResponse) => {
      console.log("SUCCESS response", resp)
      return resp.data
    })
    .catch( (err: AxiosError) => {
      console.log("ERROR err", err)
      return ['Error ' + err]
    });
}

export function* saveAsync(action: SaveAsyncAction) {
  const memo = yield call(postMemo, action); // TODO handle error
  yield put(addMemoAction(memo))
}

function postMemo(action: SaveAsyncAction) {
  return client().createMemos("/memos", {content: action.content})
    .then((resp: AxiosResponse) => {
      console.log("SUCCESS response", resp)
      return resp.data
    })
    .catch( (err: AxiosError) => {
      console.log("ERROR err", err)
      return ['Error ' + err]
    });
}
