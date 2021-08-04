let state = {
  userLevel: "",
  name: {
    eos1: "",
    eth2: "",
    bde3: "",
    eos4: ""
  },
  data: 0,
  flag: 0,
  token: ''
};
let mutations = {
  changeLevel(state, level) {
    state.userLevel = level;
  },
  changeData(state, data) {
    state.data = data;
  },
  changeName(state, name) {
    state.name = name;
  },
  changeFlag(state, flag) {
    state.flag = flag;
  },
  changeToken(state, token) {
    state.token = token;
  }
};
export default {
  state,
  mutations
};
