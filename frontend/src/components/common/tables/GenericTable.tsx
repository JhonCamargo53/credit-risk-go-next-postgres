'use client'
import React, { useEffect, useState } from 'react'
import LoadingTableItems from '../loading/LoadingTableItems'
import { TABLE_DATA_LIMIT_PER_PAGE, TABLE_DATA_LIMIT_PER_PAGE_SMALL } from '@/config/table.config';
import Button from '../buttons/Button';
import { DataColumn, GenericTableRow, PageData } from '@/types/table';

interface GenericTableProps<T extends { ID: number, UpdatedAt: string }> {
    data: T[];
    row: React.ReactElement<GenericTableRow<T>>;
    loading: boolean;
    pageData: PageData;
    columns: ReadonlyArray<DataColumn>;
    handlePagination: (pageData: PageData) => void;
    sizeVariant?: 'normal' | 'small'
}

const GenericTable = <T extends { ID: number, UpdatedAt: string }>({ data, row, loading, pageData, columns, handlePagination, sizeVariant = 'normal' }: GenericTableProps<T>) => {

    const [limit, setLimit] = useState<number>();
    const [rowLimitOptions, setRowLimitOptions] = useState<number[]>([5, 10, 20, 30]);

    useEffect(() => {
        rowLimitOptionFormat()
    }, [data])

    const rowLimitOptionFormat = () => {

        const baseOptions = sizeVariant == 'normal' ? [5, 10, 20, 30] : [5];
        let options = [...baseOptions];

        if (!options.includes(data.length) && data.length > pageData.limit) {
            options.push(data.length);
            setLimit(data.length);
        } else {
            setLimit(pageData.limit)
        }
        const newRowLimitOptions = options.sort((a, b) => a - b);
        setRowLimitOptions(newRowLimitOptions);
    }

    {/* Controlador para la siguiente página */ }
    const handleNext = () => {
        handlePagination({ ...pageData, currentPage: Math.min(pageData.currentPage + 1, pageData.lastPage) })
    }
    {/* Controlador para la página anterior */ }
    const handlePrev = () => {
        handlePagination({ ...pageData, currentPage: Math.max(pageData.currentPage - 1, 1) })
    }

    {/* Controlador de resultados por página */ }
    const handleChangeRowsPerPage = (amount: number) => {
        handlePagination({ ...pageData, currentPage: 1, limit: amount })
    };

    return (

        <div className="overflow-x-auto rounded-xl overflow-hidden border border-gray-300 text-neutral-dark">
            <table className="min-w-full table-auto">
                <thead className="text-xs uppercase bg-primary text-white text-center">
                    <tr>
                        {
                            columns.map(column => (
                                <th key={column.id} className={`px-4 py-4 text-center border-b text-base ${column.className}`} style={{ minWidth: column.minWidth ? column.minWidth : 'auto', maxWidth: column.maxWidth ? column.maxWidth : 'auto' }}>{column.label}</th>
                            ))
                        }
                    </tr>
                </thead>

                {loading ? (
                    <LoadingTableItems rowsAmount={sizeVariant == "normal" ? TABLE_DATA_LIMIT_PER_PAGE : TABLE_DATA_LIMIT_PER_PAGE_SMALL} colAmount={columns.length} />
                ) : (
                    <tbody className="bg-white divide-y divide-gray-200">

                        {data.length === 0 ? (
                            <tr className="p-3 text-center font-bold">
                                <td colSpan={columns.length} className="py-10">
                                    No hay registros de datos en el sistema
                                </td>
                            </tr>
                        ) : (
                            <>
                                {data.map((dataItem, index) => (
                                    React.cloneElement(row, {
                                        key: `${dataItem.ID}-${dataItem.UpdatedAt}`,
                                        data: dataItem,
                                        index,
                                        pageData,
                                    })
                                ))}

                                {Array.from({
                                    length: sizeVariant === 'normal'
                                        ? (pageData.limit > TABLE_DATA_LIMIT_PER_PAGE
                                            ? TABLE_DATA_LIMIT_PER_PAGE - data.length
                                            : pageData.limit - data.length)
                                        : (pageData.limit > TABLE_DATA_LIMIT_PER_PAGE_SMALL
                                            ? TABLE_DATA_LIMIT_PER_PAGE_SMALL - data.length
                                            : pageData.limit - data.length)
                                }).map((_, index) => (<tr key={`empty-${index}`} className="h-12">
                                    <td colSpan={columns.length} className={`border-b border-gray-300 hover:bg-gray-200 text-center font-medium ${(index + data.length) % 2 !== 0 ? 'bg-primary/8' : ''}`}>
                                    </td>
                                </tr>
                                ))}
                            </>
                        )}

                        <tr className='bg-neutral-dark text-white'>
                            <td colSpan={columns.length}>
                                <div
                                    className={`flex flex-row sm:items-center gap-2 sm:gap-6  font-bold p-3 
                                                ${pageData.lastPage > 1 ? 'justify-between' : 'justify-center'}
                                                   ${data.length === 0 ? 'invisible' : ''}`}
                                >
                                    {/* Selector de resultados por página */}
                                    <div className="flex items-center gap-2 text-sm">
                                        <label htmlFor="rowsPerPage">Resultados por página:</label>
                                        <select
                                            id="rowsPerPage"
                                            value={limit}
                                            onChange={(e) => handleChangeRowsPerPage(Number(e.target.value))}
                                            className="border border-gray-300 rounded px-2 py-1 text-sm text-white"
                                        >
                                            {rowLimitOptions.map((option) => (
                                                <option key={option} value={option} className='text-black'>
                                                    {option}
                                                </option>
                                            ))}
                                        </select>
                                    </div>

                                    {/* Controles de paginación */}
                                    <div className="flex items-center gap-4 justify-center">
                                        {pageData.currentPage !== 1 ? (
                                            <Button onClick={handlePrev} loading={loading}>
                                                Anterior
                                            </Button>
                                        ) : pageData.lastPage > 1 ? (
                                            <div className="w-[96px]" />
                                        ) : null}

                                        <span className="text-sm text-center">
                                            Página {pageData.currentPage} de {pageData.lastPage == 0 ? (data.length > 0 ? 1 : 0) : (pageData.lastPage)}
                                        </span>

                                        {pageData.currentPage !== pageData.lastPage && pageData.lastPage > 1 ? (
                                            <Button onClick={handleNext} loading={loading}>
                                                Siguiente
                                            </Button>
                                        ) : pageData.lastPage > 1 ? (
                                            <div className="w-[96px]" />
                                        ) : null}
                                    </div>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                )}
            </table>
        </div>
    )
}

export default GenericTable