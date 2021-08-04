<template>
  <div class="input-item">
    <div class="item-container">
      <span>{{label}}</span>
      <input :type="type" :placeholder="placeholder" :value="value" @input="inputChanged" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'input-item',
  props: ['label', 'type', 'placeholder', 'value'],
  data () {
    return {
    }
  },
  methods: {
    inputChanged (val) {
      this.$emit('update:value', val.target.value)
    }
  }
}
</script>

<style lang="scss" scoped>
.input-item {
  width: 100%;
  border-bottom: solid .1rem $divideColor;
  display: flex;
  justify-content: center;
  .item-container {
    width: 33.5rem;
    height: 8.4rem;
    display: flex;
    flex-direction: column;
    span {
      margin-top: 2.5rem;
      color: $textColor;
      font-size: 1.4rem;
    }
    input {
      padding: 1rem 0;
      color: $lightTextColor;
      font-size: 1.2rem;
    }
  }
}
</style>
