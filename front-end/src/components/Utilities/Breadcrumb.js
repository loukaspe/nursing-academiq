import React from 'react';
import { Link } from 'react-router-dom';
import "./Breadcrumb.css";

const Breadcrumb = ({ actualPath, namePath }) => {
    const pathArray = actualPath.split('/').filter(part => part);
    const nameArray = namePath.split('/').filter(name => name);

    return (
        <nav aria-label="breadcrumb">
            <ul className="breadcrumb">
                {pathArray.map((part, index) => {
                    const to = `/${pathArray.slice(0, index + 1).join('/')}`;
                    const name = nameArray[index];

                    return (
                        <li key={to} className="breadcrumb-item">
                            {index === pathArray.length - 1 ? (
                                name
                            ) : (
                                <Link to={to}>{name}</Link>
                            )}
                        </li>
                    );
                })}
            </ul>
        </nav>
    );
};

export default Breadcrumb;
