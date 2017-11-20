<template>
</template>

<script>
export default {
  methods: {
    dateToString (epoch) {
      const date = new Date(epoch)
      const toDoubleDigits = function (num) {
        num += ''
        if (num.length === 1) {
          num = '0' + num
        }
        return num
      }
      const yyyy = date.getFullYear()
      const mm = toDoubleDigits(date.getMonth() + 1)
      const dd = toDoubleDigits(date.getDate())
      const hh = toDoubleDigits(date.getHours())
      const mi = toDoubleDigits(date.getMinutes())
      const ss = toDoubleDigits(date.getSeconds())
      const sss = date.getMilliseconds()
      return `${yyyy}/${mm}/${dd} ${hh}:${mi}:${ss}.${sss} `
    }
  }
}
</script>