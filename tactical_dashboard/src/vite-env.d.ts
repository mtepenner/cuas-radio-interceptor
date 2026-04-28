/// <reference types="vite/client" />

declare module '*.glsl?raw' {
  const shader: string;
  export default shader;
}
