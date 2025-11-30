"use client";

import React, { useCallback } from "react";
import Particles from "react-tsparticles";
import type { Engine, Container } from "tsparticles-engine";
import { loadStarsPreset } from "tsparticles-preset-stars";

const StarsBackground: React.FC = () => {
  const particlesInit = useCallback(async (engine: Engine) => {
    await loadStarsPreset(engine);
  }, []);

  const particlesLoaded = useCallback(async (container: Container | undefined) => {
    
  }, []);

  return (<Particles
    id="tsparticles-stars-background"
    init={particlesInit}
    loaded={particlesLoaded}
    options={{
      preset: "stars",
      background: {
        color: {
          value: "#25282a",
        }
      },
      fullScreen: {
        enable: true,
        zIndex: -1,
      },
      particles: {
          color: {
            value: "#ffffff",
          },
        },
    }}
  />
  );
};

export default StarsBackground;
