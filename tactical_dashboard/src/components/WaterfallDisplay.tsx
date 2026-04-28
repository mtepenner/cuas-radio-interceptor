import { useEffect, useRef } from 'react';
import {
  DataTexture,
  FloatType,
  LinearFilter,
  Mesh,
  OrthographicCamera,
  PlaneGeometry,
  RedFormat,
  Scene,
  ShaderMaterial,
  WebGLRenderer,
} from 'three';

import fragmentShader from '../rendering/spectrogram_shader.glsl?raw';
import { Scan } from '../types';

const WIDTH = 64;
const HEIGHT = 48;

const vertexShader = `
varying vec2 vUv;
void main() {
  vUv = uv;
  gl_Position = vec4(position.xy, 0.0, 1.0);
}`;

type Props = {
  scan: Scan | null;
};

export function WaterfallDisplay({ scan }: Props) {
  const mountRef = useRef<HTMLDivElement | null>(null);
  const textureRef = useRef<DataTexture | null>(null);
  const rendererRef = useRef<WebGLRenderer | null>(null);
  const bufferRef = useRef<Float32Array>(new Float32Array(WIDTH * HEIGHT));
  const sceneRef = useRef<Scene | null>(null);
  const cameraRef = useRef<OrthographicCamera | null>(null);

  useEffect(() => {
    if (!mountRef.current) {
      return;
    }

    const mount = mountRef.current;
    const renderer = new WebGLRenderer({ antialias: true });
    renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2));
    renderer.setSize(mount.clientWidth, mount.clientHeight);
    mount.appendChild(renderer.domElement);

    const texture = new DataTexture(bufferRef.current, WIDTH, HEIGHT, RedFormat, FloatType);
    texture.minFilter = LinearFilter;
    texture.magFilter = LinearFilter;
    texture.needsUpdate = true;

    const material = new ShaderMaterial({
      fragmentShader,
      vertexShader,
      uniforms: {
        uTexture: { value: texture },
      },
    });

    const scene = new Scene();
    const camera = new OrthographicCamera(-1, 1, 1, -1, 0.1, 10);
    camera.position.z = 1;

    scene.add(new Mesh(new PlaneGeometry(2, 2), material));
    renderer.render(scene, camera);

    textureRef.current = texture;
    rendererRef.current = renderer;
    sceneRef.current = scene;
    cameraRef.current = camera;

    const handleResize = () => {
      if (!mountRef.current || !rendererRef.current || !sceneRef.current || !cameraRef.current) {
        return;
      }
      rendererRef.current.setSize(mountRef.current.clientWidth, mountRef.current.clientHeight);
      rendererRef.current.render(sceneRef.current, cameraRef.current);
    };

    window.addEventListener('resize', handleResize);
    return () => {
      window.removeEventListener('resize', handleResize);
      material.dispose();
      texture.dispose();
      renderer.dispose();
      mount.removeChild(renderer.domElement);
    };
  }, []);

  useEffect(() => {
    if (!scan || !textureRef.current || !rendererRef.current || !sceneRef.current || !cameraRef.current) {
      return;
    }

    const row = resample(scan.spectrum, WIDTH);
    const buffer = bufferRef.current;
    buffer.copyWithin(WIDTH, 0, WIDTH * (HEIGHT - 1));
    buffer.set(row, 0);
    textureRef.current.needsUpdate = true;
    rendererRef.current.render(sceneRef.current, cameraRef.current);
  }, [scan]);

  return <div className="waterfall-shell" ref={mountRef} />;
}

function resample(input: number[], width: number): Float32Array {
  const output = new Float32Array(width);
  if (!input.length) {
    return output;
  }

  for (let index = 0; index < width; index += 1) {
    const sampleIndex = Math.min(input.length - 1, Math.floor((index / width) * input.length));
    output[index] = normalizePower(input[sampleIndex]);
  }

  return output;
}

function normalizePower(powerDb: number): number {
  const normalized = (powerDb + 110) / 75;
  return Math.max(0, Math.min(1, normalized));
}
