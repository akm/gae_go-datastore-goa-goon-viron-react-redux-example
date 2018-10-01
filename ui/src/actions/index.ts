import {Action} from 'redux'
import { Memo } from "../states/Memo"

export enum IndexActionType {
  REFRESH_POST = "REFRESH_POST",
  ADD_MEMO = "ADD_MEMO",
  SAVE_ASYNC = "SAVE_ASYNC",
  REFRESH_ASYNC = "REFRESH_ASYNC",
}

interface RefreshPostAction extends Action {
  type: IndexActionType.REFRESH_POST;
  memos: Memo[];
}
export const refreshPostAction = (memos: Memo[]): RefreshPostAction => ({
  type: IndexActionType.REFRESH_POST,
  memos: memos
})

export interface AddMemoAction extends Action {
  type: IndexActionType.ADD_MEMO;
  memo: Memo;
}
export const addMemoAction = (memo: Memo): AddMemoAction => ({
  type: IndexActionType.ADD_MEMO,
  memo: memo
})

export interface SaveAsyncAction extends Action {
  type: IndexActionType.SAVE_ASYNC;
  content: string;
}
export const saveAsyncAction = (content: string): SaveAsyncAction => ({
  type: IndexActionType.SAVE_ASYNC,
  content: content
})

export interface RefreshAsyncAction extends Action {
  type: IndexActionType.REFRESH_ASYNC
}
export const refreshAsyncAction = (): RefreshAsyncAction => ({
  type: IndexActionType.REFRESH_ASYNC,
})

export type IndexAction = RefreshPostAction | AddMemoAction | SaveAsyncAction | RefreshAsyncAction
