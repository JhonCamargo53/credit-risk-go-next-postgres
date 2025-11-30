import { TableRow, TableCell, Skeleton, TableBody } from '@mui/material';
import React from 'react'

interface Props {
    rowsAmount: number,
    colAmount: number
}

const LoadingTableItems: React.FC<Props> = ({ rowsAmount, colAmount }) => {
    return (
        <tbody className=' divide-gray-200 bg-white'>
            {[...Array(rowsAmount)].map((_, rowIndex) => (
                <tr key={'skeleton-' + rowIndex} className={`hover:bg-gray-200   ${rowIndex % 2 !== 0 ? 'bg-primary/8' : ''}`}>
                    {[...Array(colAmount)].map((_, colIndex) => (
                        <td key={colIndex} className='px-4 py-2 text-center'>
                            <Skeleton animation="wave" variant="text" height={40} />
                        </td>
                    ))}
                </tr>
            ))}
        </tbody>
    );
}


export default LoadingTableItems