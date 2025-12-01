'use client';

import React, { createContext, useState, ReactNode, useEffect } from 'react';
import { PageData } from '@/types/table';
import { createCreditRequest, deleteCreditRequest, fetchCreditRequestById, fetchCreditRequests, fetchCreditRequestsByCustomerId, updateCreditRequest } from '@/services/creditRequest';

export interface CreditRequestContextType {
    creditRequests: CreditRequest[];
    allCreditRequests: CreditRequest[];
    selectedCreditRequest: CreditRequest | null;

    pageData: PageData;

    loading: {
        fetching: boolean;
        creating: boolean;
        updating: boolean;
        deleting: boolean;
    };

    error: string | null;
    success: string | null;

    addOrUpdateCreditRequest: (creditRequest: CreditRequest) => void;
    resetStatus: () => void;
    handlePagination: (pageData: PageData) => void;

    fetchCreditRequests: () => Promise<void>;
    fetchCreditRequestById: (id: number) => Promise<CreditRequest>;
    fetchCreditRequestsByCustomerId: (customerId: number) => Promise<void>;
    createCreditRequest: (data: CreditRequestForm) => Promise<void>;
    updateCreditRequest: (id: number, data: Partial<CreditRequest>) => Promise<void>;
    deleteCreditRequest: (id: number) => Promise<void>;
}

export const CreditRequestContext = createContext<CreditRequestContextType | undefined>(undefined);

export const CreditRequestProvider = ({ children }: { children: ReactNode }) => {
    const [creditRequests, setCreditRequests] = useState<CreditRequest[]>([]);
    const [allCreditRequests, setAllCreditRequests] = useState<CreditRequest[]>([]);
    const [selectedCreditRequest, setSelectedCreditRequest] = useState<CreditRequest | null>(null);

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

    useEffect(() => {
        fetchCreditRequestsAction();
    }, []);

    const resetStatus = () => {
        setError(null);
        setSuccess(null);
    };

    const addOrUpdateCreditRequest = (updatedCreditRequest: CreditRequest) => {
        setCreditRequests(prevCreditRequests => {

            const exists = prevCreditRequests.some(creditRequest => creditRequest.ID === updatedCreditRequest.ID);

            if (exists) {
                return prevCreditRequests.map(creditRequest =>
                    creditRequest.ID === updatedCreditRequest.ID ? updatedCreditRequest : creditRequest
                );

            } else {
                return [updatedCreditRequest, ...prevCreditRequests];
            }
        });
    };

    const handlePagination = (data: PageData, sourceData?: CreditRequest[]) => {
        const { currentPage, limit } = data;
        const all = sourceData || allCreditRequests;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        setCreditRequests(paginated);
        setPageData({
            currentPage,
            limit,
            total: all.length,
            lastPage: Math.ceil(all.length / limit),
        });
    };

    const fetchCreditRequestsAction = async () => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCreditRequests();
            setAllCreditRequests(data);

            handlePagination({ ...pageData, total: data.length }, data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener solicitudes de crédito');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const fetchCreditRequestByIdAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCreditRequestById(id);
            setSelectedCreditRequest(data);
            return data;
        } catch (err: any) {
            setError(err.message || 'Error al obtener solicitud de crédito');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const fetchCreditRequestsByCustomerIdAction = async (customerId: number) => {
        try {
            setLoading(prev => ({ ...prev, fetching: true }));
            setError(null);

            const data = await fetchCreditRequestsByCustomerId(customerId);
            setAllCreditRequests(data);

            handlePagination({ ...pageData, total: data.length }, data);
        } catch (err: any) {
            setError(err.message || 'Error al obtener solicitudes de crédito del cliente');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, fetching: false }));
        }
    };

    const createCreditRequestAction = async (data: CreditRequestForm) => {
        try {
            setLoading(prev => ({ ...prev, creating: true }));
            setError(null);
            setSuccess(null);

            const newRequest = await createCreditRequest(data);
            setAllCreditRequests(prev => {
                const newData = [newRequest, ...prev];
                handlePagination(pageData, newData);
                return newData;
            });

            setSuccess('Solicitud de crédito creada correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al crear solicitud de crédito');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, creating: false }));
        }
    };

    const updateCreditRequestAction = async (id: number, data: Partial<CreditRequest>) => {
        try {
            setLoading(prev => ({ ...prev, updating: true }));
            setError(null);
            setSuccess(null);

            const updatedRequest = await updateCreditRequest(id, data);

            setAllCreditRequests(prev => {
                const newData = prev.some(creditRequest => creditRequest.ID === id)
                    ? prev.map(creditRequest => (creditRequest.ID === id ? updatedRequest : creditRequest))
                    : [updatedRequest, ...prev];
                handlePagination(pageData, newData);
                console.log(newData)
                return newData;
            });

            setSelectedCreditRequest(updatedRequest);
            setSuccess('Solicitud de crédito actualizada correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al actualizar solicitud de crédito');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, updating: false }));
        }
    };

    const deleteCreditRequestAction = async (id: number) => {
        try {
            setLoading(prev => ({ ...prev, deleting: true }));
            setError(null);
            setSuccess(null);

            await deleteCreditRequest(id);
            const updatedAll = allCreditRequests.filter(creditRequest => creditRequest.ID !== id);
            setAllCreditRequests(updatedAll);
            handlePagination(pageData, updatedAll);
            setSuccess('Solicitud de crédito eliminada correctamente');
        } catch (err: any) {
            setError(err.message || 'Error al eliminar solicitud de crédito');
            throw err;
        } finally {
            setLoading(prev => ({ ...prev, deleting: false }));
        }
    };

    return (
        <CreditRequestContext.Provider value={{
            creditRequests,
            allCreditRequests,
            selectedCreditRequest,
            pageData,
            loading,
            error,
            success,
            resetStatus,
            addOrUpdateCreditRequest,
            handlePagination,
            fetchCreditRequests: fetchCreditRequestsAction,
            fetchCreditRequestById: fetchCreditRequestByIdAction,
            fetchCreditRequestsByCustomerId: fetchCreditRequestsByCustomerIdAction,
            createCreditRequest: createCreditRequestAction,
            updateCreditRequest: updateCreditRequestAction,
            deleteCreditRequest: deleteCreditRequestAction,
        }}>
            {children}
        </CreditRequestContext.Provider>
    );
};
