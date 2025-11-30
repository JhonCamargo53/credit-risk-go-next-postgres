import { ModalContext, ModalContextType } from '@/context/ModalContext';
import { useContext } from 'react';

export const useModal = (): ModalContextType => {
    const context = useContext(ModalContext);
    if (context === undefined) {
        throw new Error('useModal debe usarse dentro de un ModalProvider');
    }
    return context;
};
