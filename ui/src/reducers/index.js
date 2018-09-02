export default (state = [], action) => {
  console.log("reducer state", state)
  console.log("reducer action", action)

  switch (action.type) {
  case 'REFRESH_POST':
    return action.memos
  case 'ADD_MEMO':
    return [...state, action.memo]
  default:
    return state
  }
}
