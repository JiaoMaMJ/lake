module.exports = {
  // Configuration of lake's own services
  lake: {
    // Enable basic authentication to the lake API
    // token: 'mytoken'
    // This is the time between fetching new data from all sources on automated loop
    loopIntervalInMinutes: 60
  },
  // Configuration of MongoDB
  mongo: {
    connectionString: 'mongodb://lake:lakeIScoming@mongodb:27017/lake?authSource=admin'
  },
  // Configuration of rabbitMQ
  rabbitMQ: {
    connectionString: 'amqp://guest:guestWhat@rabbitmq:5672/rabbitmq'
  },
  // Configuration of PostgreSQL
  postgres: {
    username: 'postgres',
    password: 'postgresWhat',
    host: 'postgresdb',
    database: 'lake',
    port: 5432,
    dialect: 'postgres'
  },
  // Configuration of Jira plugin
  jira: {
    // Replace example host with your own host
    host: 'https://your-domain.atlassian.net',
    // Replace *** with your jira API token, please see Jira plugin readme for details
    basicAuth: '***',
    // Enable proxy for interacting with Jira API
    // proxy: 'http://localhost:4780',
    // Set timeout for sending requests to Jira API
    timeout: 10000,
    // Set max retry times for sending requests to Jira API
    maxRetry: 3
    // This property is used to automatically collect data from jira every 60 minutes
    // dataCollection: {
    //   "boardId": 0
    // }
  },
  // Configuration of Gitlab plugin
  gitlab: {
    // Replace example host with your host if your Gitlab is self-hosted
    // Leave this unchanged if you use Gitlab's cloud service
    host: 'https://gitlab.com',
    apiPath: 'api/v4',
    // Replace *** with your Gitlab API token, please see Gitlab plugin readme for details
    token: '***',
    // Enable proxy for interacting with Gitlab API
    // proxy: 'http://localhost:4780',
    // Set timeout for sending requests to Jira API
    timeout: 10000,
    // Set max retry times for sending requests to Jira API
    maxRetry: 3,
    // Customize what data you'd like to collect
    // Leave this unchanged if you wish to collect everything
    skip: {
      commits: false,
      projects: false,
      mergeRequests: false,
      notes: false
    }
    // This property is used to automatically collect data from gitlab every 60 minutes
    // dataCollection: {
    //   "projectId": 0
    // }
  },
  // COnfiguration of Jira <> Gitlab mapping
  jiraBoardGitlabProject: {
    // Replace example mapping your own mapping
    // Format: <Jira boardID>: <Gitlab projectId>
    8: 8967944
  }
}
