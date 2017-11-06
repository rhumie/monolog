<template>
<div>
  <div class="container">
    <div role="progressbar" class="mdc-linear-progress mdc-linear-progress--indeterminate" v-show="isLoading">
      <div class="mdc-linear-progress__buffering-dots"></div>
      <div class="mdc-linear-progress__buffer"></div>
      <div class="mdc-linear-progress__bar mdc-linear-progress__primary-bar">
        <span class="mdc-linear-progress__bar-inner"></span>
      </div>
      <div class="mdc-linear-progress__bar mdc-linear-progress__secondary-bar">
        <span class="mdc-linear-progress__bar-inner"></span>
      </div>
    </div>
    <div class="mdc-layout-grid">
      <div class="mdc-layout-grid__inner">
        <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-12">
          <div class="mdc-textfield mdc-textfield--fullwidth" :class="{'disabled': isLoading}">
            <input type="text" id="id-textfield" class="mdc-textfield__input" @keyup.enter="authenticate" v-model="id">
            <label class="mdc-textfield__label" for="id-textfield">ID</label>
            <div class="mdc-textfield__bottom-line"></div>
          </div>
        </div>
        <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-12">
          <div class="mdc-textfield mdc-textfield--fullwidth" :class="{'disabled': isLoading}">
            <input type="password" id="secret-textfield" class="mdc-textfield__input" @keyup.enter="authenticate" v-model="secret">
            <label class="mdc-textfield__label" for="secret-textfield">Secret</label>
            <div class="mdc-textfield__bottom-line"></div>
          </div>
        </div>
      </div>
      <div class="message-area" v-show="hasError">
        <span class="error">Incorrect id or secret. Please try again.</span>
      </div>
      <div class="button-area right-align" :class="{'disabled': isLoading}">
      <button class="mdc-button mdc-button--unelevated" @click="authenticate">Login</button>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import AWS from 'aws-sdk'
import { MDCTextfield } from '@material/textfield'

export default {
  name: 'Login',
  data () {
    return {
      id: '',
      secret: '',
      hasError: false,
      isLoading: false
    }
  },
  methods: {
    authenticate: function () {
      this.hasError = false
      this.isLoading = true
      var sts = new AWS.STS(new AWS.Config({
        accessKeyId: this.id,
        secretAccessKey: this.secret
      }))
      var self = this
      sts.getSessionToken({DurationSeconds: 3600}, (err, data) => {
        if (err) {
          self.hasError = true
          this.isLoading = false
        } else {
          const credential = data.Credentials
          localStorage.setItem('token', `${credential.AccessKeyId}:${credential.SecretAccessKey}:${credential.SessionToken}`)
          self.$router.push('/#', () => { this.isLoading = false })
        }
      })
    }
  },
  mounted () {
    this.$el.querySelectorAll('.mdc-textfield').forEach((element) => {
      MDCTextfield.attachTo(element)
    })
  }
}
</script>

<style scoped>
.container {
  max-width: 500px;
  margin-top: 100px;
  background-color: #EEEEEE;
}

.mdc-textfield {
  height: 48px;
  font-size: 14px;
}

.message-area {
  margin: 45px 0
}

.error {
  color: #F44336
}

.button-area {
  margin-top: 45px
}
</style>
