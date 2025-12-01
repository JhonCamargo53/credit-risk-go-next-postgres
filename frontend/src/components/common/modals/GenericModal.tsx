'use client';
import React, { ReactNode, useEffect, useState, useCallback } from 'react';
import { createPortal } from 'react-dom';
import { IoClose } from 'react-icons/io5';
import clsx from 'clsx';
import Button from '../buttons/Button';
import { useModal } from '@/hooks/useModal';

interface GenericModalProps {
  title?: string;
  isOpen: boolean;
  setOpen: (state: boolean) => void;
  content: ReactNode;
  zIndex?: number;
  closeIcon?: boolean;
  actionButton?: ReactNode;
  size?: '2xs' | 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | '3xl' | '4xl' | '5xl' | '6xl' | '7xl' | 'full';
  handleOnClose?: () => void;
}

const widthClasses = {
  '2xs': 'md:w-48',
  'xs': 'md:w-64',
  'sm': 'md:w-80',
  'md': 'md:w-[28rem]',
  'lg': 'md:w-[36rem]',
  'xl': 'md:w-[42rem]',
  '2xl': 'md:w-[48rem]',
  '3xl': 'md:w-[56rem]',
  '4xl': 'md:w-[64rem]',
  '5xl': 'md:w-[72rem]',
  '6xl': 'md:w-full',
  '7xl': 'md:w-full',
  'full': 'md:w-full'
};

const GenericModal: React.FC<GenericModalProps> = ({
  title,
  isOpen,
  setOpen,
  content,
  zIndex = 100,
  closeIcon = true,
  actionButton,
  handleOnClose = () => { },
  size = 'md',
}) => {

  const { addOpenModal, openModalList, removeLastOpenModal } = useModal();
  const [modalId] = useState(`modal-${new Date().toUTCString()}`);
  const [animateIn, setAnimateIn] = useState(false);
  const [shouldRender, setShouldRender] = useState(false);

  useEffect(() => {
    if (isOpen) {

      setShouldRender(true);
      document.body.style.overflow = 'hidden';
      setTimeout(() => setAnimateIn(true), 10);
      addOpenModal(modalId);

    } else if (shouldRender) {
      setAnimateIn(false);
      setTimeout(() => setShouldRender(false), 300);
    }
  }, [isOpen]);

  useEffect(() => {
    const currentExist = openModalList.includes(modalId);
    if (shouldRender && !currentExist) {
      setShouldRender(false);
      setOpen(false);
      if (openModalList.length === 0) {
        document.body.style.overflow = '';
      }
    }
  }, [openModalList, shouldRender]);

  const handleKeyDown = useCallback((e: KeyboardEvent) => {
    if (e.key === 'Escape') {
      if (openModalList[openModalList.length - 1] === modalId) {
        document.body.style.overflow = '';
        handleOverlayClick();
      }
    }
  }, [openModalList, modalId]);

  useEffect(() => {

    if (shouldRender) {
      window.addEventListener('keydown', handleKeyDown);
    }

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };

  }, [shouldRender, handleKeyDown]);

  const handleOverlayClick = () => {
    document.body.style.overflow = '';

    setAnimateIn(false);
    setTimeout(() => {
      handleOnClose()
      removeLastOpenModal()
    }, 300);
  };

  if (!shouldRender) return null;

  return createPortal(
    <div className={`fixed inset-0 z-[${zIndex}]`}>

      {/* Overlay */}
      <div
        className="absolute inset-0 bg-black/50 backdrop-blur-sm"
        onClick={handleOverlayClick}
      />

      {/* Modal Panel */}
      <div
        className={clsx(
          'fixed inset-0 bg-neutral-light shadow-lg transition-transform duration-300 ease-in-out flex flex-col',
          'w-full h-full',
          widthClasses[size],
          {
            'translate-x-0': animateIn,
            'translate-x-full md:translate-x-full': !animateIn,
            'md:top-0 md:right-0 md:bottom-0 md:left-auto md:h-screen': true,
          }
        )}
      >
        <div className="bg-neutral-dark h-[60px] px-4 w-full" />

        {/* Scrollable content */}
        <div className="px-6 py-4 w-full flex-1 min-h-0 overflow-auto space-y-2 scrollbar-primary">
          <div className={`w-full flex ${title ? 'justify-between ' : 'justify-end'}`}>
            {title && <h2 className="text-lg font-semibold text-primary uppercase">{title}</h2>}
            {closeIcon && (
              <button onClick={handleOverlayClick} className="text-black hover:text-error transition cursor-pointer">
                <IoClose size={24} />
              </button>
            )}
          </div>
          {content}
        </div>
      </div>
    </div>,
    document.body
  );
};

export default GenericModal;
