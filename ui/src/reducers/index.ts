import { IndexActionType, IndexAction } from '../actions/index'

export default (state = [], action: IndexAction) => {
  console.log("reducer state", state)
  console.log("reducer action", action)

  switch (action.type) {
  case IndexActionType.REFRESH_POST:
    return action.memos
  case IndexActionType.ADD_MEMO:
    return [...state, action.memo]
  default:
    return state
  }
}
