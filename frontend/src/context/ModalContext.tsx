'use client'
import React, { createContext, ReactNode, useEffect, useState } from 'react';

export interface ModalContextType {
    openModalList: string[];
    addOpenModal: (modalId: string) => void;
    removeOpenModal: (modalId: string) => void;
    removeLastOpenModal: () => void
}

export const ModalContext = createContext<ModalContextType | undefined>(undefined);

export const ModalProvider = ({ children }: { children: ReactNode }) => {

    const [openModalList, setOpenModalList] = useState<string[]>([]);

    const addOpenModal = (modalId: string) => {
        setOpenModalList([...openModalList, modalId]);
    }

    const removeOpenModal = (modalId: string) => {
        setOpenModalList(openModalList.filter(modal => modal !== modalId));
    }

    const removeLastOpenModal = () => {
        setOpenModalList(prevList => prevList.slice(0, -1));
    }

    useEffect(() => {
        if (openModalList.length === 0) {
            document.body.style.overflow = '';
        }
    }, [openModalList]);

    return (
        <ModalContext.Provider value={{
            openModalList,
            addOpenModal,
            removeOpenModal,
            removeLastOpenModal
        }}>
            {children}
        </ModalContext.Provider>
    );
};
