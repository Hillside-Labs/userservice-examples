import React, { useState, useEffect } from 'react';
import axios from 'axios';

/* 
    Example embedding a form for issuing a team invite.
    /api/teams/invites/new/form returns html with the form.
    You can always send your own form and POST it to /api/teams/invites/new/form.
 */
function CardWithForm({baseURL}) {
    const [htmlContent, setHtmlContent] = useState('');
    userId = session.GetUserID() //ID of the user who issued the invite, so we can attach the 'inviter' attribute to them
    redirect = "/"               //This is where the user gets redirected after issuing the invite!
    useEffect(() => {

        axios.get(`${baseURL}/api/teams/invites/new/form?userid=${userId}&redirect=${redirect}`)
            .then(response => {
                setHtmlContent(response.data);
            })
            .catch(error => {
                console.error('Error fetching HTML content:', error);
            });
    }, []);

    return (
        <div className="card">
            <div className="card-body">
                {/* Render the HTML content */}
                <div dangerouslySetInnerHTML={{ __html: htmlContent }} />
            </div>
        </div>
    );
}

export default CardWithForm;