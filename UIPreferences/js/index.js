const { Configuration, DefaultApi } = require('userup-sdk')
require('dotenv').config()

async function main() {
    const userid = parseInt(process.env.USERID, 10)
    
    const API_BASE_PATH = 'http://localhost:9001/api'
    
    const config = new Configuration({ basePath: API_BASE_PATH })
    const api = new DefaultApi(config)
    
    try {
        const attrReq = {
            id: userid,
            requestBody: {
                'uiPreferences': {
                    theme: 'darcula',
                    fontSize: 12,
                    language: 'en-IN',
                    showNotifications: true,
                }
            }
        }
        await api.createAttributes(attrReq)
        
        console.log(await api.getAttributes({ id: userid }))

    } catch (error) {
        console.error(error)
    }
}

main().catch(console.error)
