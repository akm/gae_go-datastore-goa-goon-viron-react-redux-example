import { call, takeEvery, put } from 'redux-saga/effects'

import { MemosApiFactory, Memo, MemoCollection } from '../api/'
import { IndexActionType, SaveAsyncAction, addMemoAction, refreshPostAction } from '../actions/'

export default function* rootSaga() {
  yield takeEvery(IndexActionType.REFRESH_ASYNC, refreshAsync)
  yield takeEvery(IndexActionType.SAVE_ASYNC, saveAsync)
}

export function* refreshAsync() {
  const memos = yield call(fetchMemos); // TODO handle error
  yield put(refreshPostAction(memos))
}

function fetchMemos() {
  return MemosApiFactory().memosList()
    .then((memos: MemoCollection) => {
      console.log("SUCCESS response", memos);
      return memos;
    })
    .catch((err) => {
      console.log("ERROR", err)
      return err
    });
}

export function* saveAsync(action: SaveAsyncAction) {
  const memo = yield call(postMemo, action); // TODO handle error
  yield put(addMemoAction(memo))
}

function postMemo(action: SaveAsyncAction) {
  return MemosApiFactory().memosCreate({content: action.content})
    .then((memo: Memo) => {
      console.log("SUCCESS response", memo)
      return memo;
    })
    .catch((err) => {
      console.log("ERROR ", err);
      return err;
    })
}
