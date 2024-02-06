/* 
    Example using our smaple form for creating a team and embedding it.
    /api/teams/new/form returns html with the form.
    You can always send your own form and POST it to /api/teams/new/form.
 */
function CardWithForm({baseURL}) {
    const [htmlContent, setHtmlContent] = useState('');
    userId = session.getUserID() //If you want the user joining this team at creation, then add this field, leave it empty otherwise.
    redirect = "/"               //This is where user gets redirected after creating the team!
    useEffect(() => {

        axios.get(`${baseURL}/api/teams/new/form?userid=${userId}&redirect=${redirect}`)
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
    

/*
    Example using your own form to create a team.
*/
function TeamForm({ baseURL }) {
    const [team, setTeam] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        const user_id = session.getUserID(); // Assuming session is defined elsewhere
        const redirect = '/dashboard'; // Set the redirect URL

        try {
            // Send POST request
            await axios.post(`${baseURL}/teams/new/form`, {
                user_id,
                team,
                redirect
            });

        } catch (error) {
            console.error('Error submitting form:', error);
            // Handle error
        }
    };

    return (
        <div>
            <h2>Create New Team</h2>
            <form onSubmit={handleSubmit}>
                <input type="hidden" name="user_id" value={session.getUserId()} />
                <input type="hidden" name="redirect" value={redirect} />
                <div>
                    <label htmlFor="team">Team Name:</label>
                    <input 
                        type="text" 
                        id="team" 
                        name="team" 
                        value={team} 
                        onChange={(e) => setTeam(e.target.value)} 
                        required 
                    />
                </div>
                <button type="submit">Create Team</button>
            </form>
        </div>
    );
}