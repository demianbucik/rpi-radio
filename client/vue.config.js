const dotenv = require('dotenv-webpack');

module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  configureWebpack: {
    plugins: [
      new dotenv()
    ]
  }
}