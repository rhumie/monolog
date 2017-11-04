<template>
<div>
  <div class="mdc-select" :class="{fullwidth: fullwidth}" role="listbox" tabindex="0">
    <span class="mdc-select__selected-text">{{selectedText}}</span>
    <div class="mdc-simple-menu mdc-select__menu">
      <ul class="mdc-list mdc-simple-menu__items">
        <li class="mdc-list-item" role="option" aria-disabled="true">{{selectedText}}</li>
        <li class="mdc-list-item" role="option" v-for="option in options" :id="option.value" :aria-selected="option.selected" tabindex="0">{{option.label}}</li>
      </ul>
    </div>
  </div>
  <label class="label" :for="id">{{label}}</label>
</div>
</template>

<script>
import { MDCSelect } from '@material/select'

let mdcSelect

export default {
  name: 'select-menu',
  props: {
    id: {
      type: String,
      required: true
    },
    label: {
      type: String
    },
    options: {
      type: Array,
      required: true
    },
    fullwidth: {
      type: Boolean
    }
  },
  data () {
    return {
      state: {
        value: null
      }
    }
  },
  mounted () {
    mdcSelect = MDCSelect.attachTo(this.$el.querySelector('.mdc-select'))
    mdcSelect.listen('MDCSelect:change', () => {
      this.state.value = mdcSelect.value
      this.$emit('change', mdcSelect.value)
    })
  },
  computed: {
    selectedText () {
      return 'Select ' + this.label
    }
  }
}
</script>

<style scoped>
.fullwidth {
  width: 100% !important
}

.label {
  color: var(--mdc-theme-text-hint-on-light, rgba(0, 0, 0, 0.5));
  display: block;
  transform-origin: left top;
  transform: translateY(-250%) scale(0.75, 0.75)
}

.mdc-select {
  height: 48px;
}
</style>
