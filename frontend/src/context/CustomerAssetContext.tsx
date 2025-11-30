'use client';

import React, { createContext, useState, ReactNode } from 'react';
import { PageData } from '@/types/table';
import {
    fetchCustomerAssets,
    fetchCustomerAssetById,
    createCustomerAsset,
    updateCustomerAsset,
    deleteCustomerAsset,
    fetchCustomerAssetsByCreditRequest
} from '@/services/customerAssetService';
import { CustomerAsset, CustomerAssetForm } from '@/types/customerAsset';

export interface CustomerAssetContextType {
    customerAssets: CustomerAsset[];
    allCustomerAssets: CustomerAsset[];
    selectedCustomerAsset: CustomerAsset | null;

    pageData: PageData;

    loading: {
        fetching: boolean;
        creating: boolean;
        updating: boolean;
        deleting: boolean;
    };

    error: string | null;
    success: string | null;

    resetStatus: () => void;
    handlePagination: (pageData: PageData) => void;

    fetchCustomerAssets: () => Promise<void>;
    fetchCustomerAssetById: (id: number) => Promise<void>;
    fetchCustomerAssetsByCreditRequestId: (id: number) => Promise<void>;
    createCustomerAsset: (data: CustomerAssetForm) => Promise<void>;
    updateCustomerAsset: (id: number, data: Partial<CustomerAsset>) => Promise<void>;
    deleteCustomerAsset: (id: number) => Promise<void>;
}

export const CustomerAssetContext = createContext<CustomerAssetContextType | undefined>(undefined);

export const CustomerAssetProvider = ({ children }: { children: ReactNode }) => {
    const [customerAssets, setCustomerAssets] = useState<CustomerAsset[]>([]);
    const [allCustomerAssets, setAllCustomerAssets] = useState<CustomerAsset[]>([]);
    const [selectedCustomerAsset, setSelectedCustomerAsset] = useState<CustomerAsset | null>(null);

    const [pageData, setPageData] = useState<PageData>({
        currentPage: 1,
        limit: 10,
        total: 0,
        lastPage: 1,
    });

    const [loading, setLoading] = useState({
        fetching: true,
        creating: false,
        updating: false,
        deleting: false,
    });

    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);

    const resetStatus = () => {
        setError(null);
        setSuccess(null);
    };

    const handlePagination = (data: PageData, sourceAssets?: CustomerAsset[]) => {
        const { currentPage, limit } = data;
        const all = sourceAssets || allCustomerAssets;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        setCustomerAssets(paginated);
        setPageData({
            currentPage,
            limit,
            total: all.length,
            lastPage: Math.ceil(all.length / limit),
        });
    };

    const fetchCustomerAssetsAction = async () => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCustomerAssets();
            setAllCustomerAssets(data);

            handlePagination({ ...pageData, total: data.length }, data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener activos de clientes');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const fetchCustomerAssetByIdAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCustomerAssetById(id);
            setSelectedCustomerAsset(data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener activo de cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const fetchCustomerAssetsByCreditRequestIdAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCustomerAssetsByCreditRequest(id);

            setAllCustomerAssets(data);
            handlePagination({ ...pageData, total: data.length }, data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener activo de cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };


    const createCustomerAssetAction = async (data: CustomerAssetForm) => {
        try {
            setLoading(prev => ({ ...prev, creating: true }));
            setError(null);
            setSuccess(null);

            const newAsset = await createCustomerAsset(data);
            setAllCustomerAssets(prev => {
                const newData = [newAsset, ...prev]
                handlePagination(pageData, newData);
                return newData
            });

            setSuccess('Activo de cliente creado correctamente');

        } catch (err: any) {
            setError(err.message || 'Error al crear activo de cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, creating: false }));
        }
    };

    const updateCustomerAssetAction = async (id: number, data: Partial<CustomerAsset>) => {
        try {
            setLoading(prev => ({ ...prev, updating: true }));
            setError(null);
            setSuccess(null);

            const updatedAsset = await updateCustomerAsset(id, data);

            setAllCustomerAssets(prev => {
                const newData = prev.some(asset => asset.ID === id)
                    ? prev.map(asset => (asset.ID === id ? updatedAsset : asset))
                    : [updatedAsset, ...prev];
                handlePagination(pageData, newData);
                return newData;
            });

            setSelectedCustomerAsset(updatedAsset);
            setSuccess('Activo de cliente actualizado correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al actualizar activo de cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, updating: false }));
        }
    };

    const deleteCustomerAssetAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, deleting: true }));
            setError(null);
            setSuccess(null);

            await deleteCustomerAsset(id);
            const updatedAll = allCustomerAssets.filter(a => a.ID !== id);
            setAllCustomerAssets(updatedAll);
            handlePagination(pageData, updatedAll);
            setSuccess('Activo de cliente eliminado correctamente');
            
        } catch (err: any) {
            setError(err.message || 'Error al eliminar activo de cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, deleting: false }));
        }
    };

    return (
        <CustomerAssetContext.Provider value={{
            customerAssets,
            allCustomerAssets,
            selectedCustomerAsset,
            pageData,
            loading,
            error,
            success,
            resetStatus,
            handlePagination,
            fetchCustomerAssets: fetchCustomerAssetsAction,
            fetchCustomerAssetsByCreditRequestId: fetchCustomerAssetsByCreditRequestIdAction,
            fetchCustomerAssetById: fetchCustomerAssetByIdAction,
            createCustomerAsset: createCustomerAssetAction,
            updateCustomerAsset: updateCustomerAssetAction,
            deleteCustomerAsset: deleteCustomerAssetAction,
        }}>
            {children}
        </CustomerAssetContext.Provider>
    );
};
