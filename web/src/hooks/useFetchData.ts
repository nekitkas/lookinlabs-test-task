import { useState, useEffect } from 'react';

export const useFetchData = <T,>(url: string) => {
    const [data, setData] = useState<T | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url);
                const data: T = await response.json();
                setData(data);
            } catch (error) {
                setError(`An error occurred: ${error}`);
            } finally {
                setLoading(false);
            }
        }
    fetchData();
    }, [url]);

    return { data, loading, error };
}