import { useState, useEffect } from 'react';

export const useFetchData = <T,>(url: string) => {
    const [data, setData] = useState<T | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url);
                if (!response.ok) {
                    throw new Error(`${response.statusText}`);
                }
                const data: T = await response.json();
                setData(data);
                setError(null);
            } catch (error) {
                setError(`${error}`);
            } finally {
                setLoading(false);
            }
        }
    fetchData();
    }, [url]);

    return { data, loading, error };
}