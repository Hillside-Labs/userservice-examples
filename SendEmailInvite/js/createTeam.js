const axios = require('axios');

async function submitTeamForm() {
    const baseURL = 'https://service.com'; // Replace with your base URL
    const user_id = '1';                
    const team = 'cool folks';

    try {
        const response = await axios.post(`${baseURL}/teams/new/form`, {
            user_id,
            team,
            redirect:"" //When the redirect value is empty, the handler will return a 201
        });


        if (response.status === 201) {
            console.log('Form submitted successfully');
            console.log('Response:', response.data); //this should just print a success message!
        } else {
            console.error('Error: Unexpected response status:', response.status);
            //Handle a bad response
        }
    } catch (error) {
        console.error('Error submitting form:', error.response ? error.response.data : error.message);
        // Handle error
    }
}
