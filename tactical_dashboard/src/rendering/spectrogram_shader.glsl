varying vec2 vUv;
uniform sampler2D uTexture;

vec3 palette(float t) {
    vec3 low = vec3(0.03, 0.11, 0.15);
    vec3 mid = vec3(0.01, 0.56, 0.43);
    vec3 hot = vec3(0.95, 0.62, 0.11);
    vec3 critical = vec3(0.92, 0.17, 0.13);
    if (t < 0.45) {
        return mix(low, mid, smoothstep(0.0, 0.45, t));
    }
    if (t < 0.8) {
        return mix(mid, hot, smoothstep(0.45, 0.8, t));
    }
    return mix(hot, critical, smoothstep(0.8, 1.0, t));
}

void main() {
    float intensity = texture2D(uTexture, vec2(vUv.x, 1.0 - vUv.y)).r;
    gl_FragColor = vec4(palette(intensity), 1.0);
}
