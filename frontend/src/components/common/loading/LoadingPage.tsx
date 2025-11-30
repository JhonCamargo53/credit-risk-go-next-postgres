'use client';

import React, { useEffect, useState } from 'react';
import { DotLoader } from 'react-spinners';

const LoadingPage = ({ loadingText = 'Cargando' }: { loadingText?: string }) => {
  const [dots, setDots] = useState<string>('...');
  const [adding, setAdding] = useState<boolean>(true);

  useEffect(() => {
    const interval = setInterval(() => {
      setDots(prev => {
        if (adding) {
          if (prev.length >= 3) {
            setAdding(false);
            return prev.slice(0, -1);
          } else {
            return prev + '.';
          }
        } else {
          if (prev.length <= 1) {
            setAdding(true);
            return '';
          } else {
            return prev.slice(0, -1);
          }
        }
      });
    }, 500);

    return () => clearInterval(interval);
  }, [adding]);

  return (
    <div className="bg-neutral-dark shadow rounded p-4 min-h-screen flex justify-center items-center">
      <div className="max-w-2xl w-full text-center p-8">
        <div className="mb-3 flex justify-center">
          <DotLoader size={150} color="white" />
        </div>
        <div className="text-5xl font-bold text-white mt-30">
          {loadingText}
          {dots}
        </div>
      </div>
    </div>
  );
};

export default LoadingPage;
