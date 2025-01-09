import React, { useEffect } from 'react';

const Download: React.FC = () => {
    useEffect(() => {
        const link = document.createElement('a');
        link.href = 'client.apk';
        link.download = 'client.apk';
        document.body.appendChild(link);
        link.click();
    }, []);
    
    return null;
};

export default Download;
