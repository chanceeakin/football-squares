import babel from "rollup-plugin-babel";
import uglify from "rollup-plugin-uglify";

export default {
  input: "./target/deploy/client.js",
  output: {
    name: "client",
    file: "./release/client.js",
    format: "es"
  },
  plugins: [
    babel({
      exclude: "node_modules/**"
    }),
    uglify
  ]
};
