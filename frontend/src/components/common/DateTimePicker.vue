<template>
<div>
  <div class="mdc-textfield mdc-textfield--fullwidth">
    <date-picker class="mdc-textfield__input" v-model="state.value"
      :id="id"
      :type="type"
      :disabled="disabled"
      :readonly="readonly"
      :editable="editable"
      :clearable="clearable"
      :startPlaceholder="startPlaceholder"
      :endPlaceholder="endPlaceholder"
      @change="updateValue"></date-picker>
      <label class="label" :for="id"><span class="label-text">{{label}}</span></label>
  </div>
</div>
</template>

<script>
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import { DatePicker } from 'element-ui'
import 'element-ui/lib/theme-chalk/date-picker.css'
import 'element-ui/lib/theme-chalk/button.css'
import 'element-ui/lib/theme-chalk/icon.css'

// Set Element locale.
locale.use(lang)

export default {
  name: 'datetime-picker',
  components: {
    datePicker: DatePicker
  },
  props: {
    id: {
      type: String
    },
    value: {
      type: Date
    },
    label: {
      type: String
    },
    readonly: {
      type: Boolean
    },
    disabled: {
      type: Boolean
    },
    editable: {
      type: Boolean
    },
    clearable: {
      type: Boolean
    },
    startPlaceholder: {
      type: String
    },
    endPlaceholder: {
      type: String
    },
    type: {
      type: String,
      default: 'datetime'
    }
  },
  data () {
    return {
      state: {
        value: this.value
      }
    }
  },
  watch: {
    value: function (value) {
      if (this.state.value !== value) {
        this.state.value = value
      }
    }
  },
  methods: {
    updateValue: function (val) {
      this.$emit('input', val)
    }
  }
}
</script>

<style>
.mdc-textfield .el-input {
  font-size: 0.936rem;
}
.mdc-textfield .el-input__inner {
  border: none;
  background-color: rgba(0,0,0,0);
  color: inherit;
  font-family: inherit;
  letter-spacing: 0.04em;
  height: 48px;
  padding: 0;
}
.mdc-textfield .el-input__inner::placeholder {
  color: var(--mdc-theme-text-hint-on-light, rgba(0, 0, 0, 0.38));
}
.mdc-textfield .el-input__inner::-webkit-input-placeholder {
  color: var(--mdc-theme-text-hint-on-light, rgba(0, 0, 0, 0.38));
}
.mdc-textfield .el-input__prefix {
  display: none;
}
.el-picker-panel .el-date-table td.current:not(.disabled) span {
  background-color: var(--mdc-theme-primary, #CDDC39);
}
.el-picker-panel .el-date-table td.today span {
  color: var(--mdc-theme-primary, #CDDC39);
}
.el-picker-panel .el-time-panel__btn.confirm {
  color: var(--mdc-theme-primary, #CDDC39);
}
.el-picker-panel .el-button--text {
  color: var(--mdc-theme-primary, #CDDC39);
}
.el-picker-panel .el-input__inner {
  border: none;
}
</style>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.mdc-textfield {
  height: 48px;
}

.label {
  color: var(--mdc-theme-text-hint-on-light, rgba(0, 0, 0, 38));
  display: block;
  transform-origin: left top;
  transform: translateY(-250%) scale(0.75, 0.75);
  white-space: nowrap;
}
</style>