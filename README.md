# GO BACKEND CHALLENGE API MICROSERVICE

## 1. Get started

```bash
  create .env file in the root of the project with the following variables

    DB_HOST=
    DB_PORT=
    DB_NAME=
    DB_USER=
    DB_PASS=
    DB_ENGINE=
    BASIC_AUTH_USERNAME=
    BASIC_AUTH_PASSWORD=
    SERVER_PORT=
    CORS_WHITELIST=
    SHOW_GORM_LOGS=
    APPLY_MIGRATIONS=  
  
```

## 2. DB Diagram

```txt
https://drive.google.com/file/d/1Knr1th8bGDwQvMLMgpX_F41EuMWsMxwC/view?usp=sharing
```

## 3. Postman

```txt
https://www.postman.com/blind-creator/workspace/github.com/luisnquin/blind-creator/test-core
```

## 4. Challenges

- [x] Add metadata

  We want to add an campaign_creator_social_network_actions array for each campaign in the list campaign response data

  ```json
  {
      "current_page": 1,
      "data": [
          {
              "campaign_id": 2,
              "campaign_created_at": "2023-05-03T17:29:25.534412-04:00",
              "campaign_updated_at": "2023-05-03T17:29:25.534412-04:00",
              "campaign_name": "test",
              "campaign_initial_date": "2023-04-20T00:00:00Z",
              "campaign_final_date": "2023-04-21T00:00:00Z",
              "campaign_budget": 100,
              "campaign_currency": "USD",
              "campaign_agency_id": 1,
              "campaign_manager_id": 1,
              "campaign_company_id": 1,
              "campaign_bundle_id": 0,
              "manager_email": "cristian.mancilla96@gmail.com",
              "manager_name": "Cristian",
              "company_name": "Company Demo",
              "company_email": "demo@demo.com",
              "campaign_creator_social_network_actions": [
                  {
                      "action_code_name": "",
                      "action_quantity": 0,
                      "action_cost_price": 0,
                      "action_cost_currency": "",
                      "action_bundle_price": 0,
                      "action_accepted_price": 0,
                      "action_draft_content_status": "",
                      "action_final_content_status": "",
                      "action_upload_draft_content_date": "",
                      "action_upload_final_content_date": "",
                      "action_payment_condition": "",
                      "action_creator_id": 0,
                      "action_creator_name": "",
                      "action_creator_avatar": "",
                      "action_creator_email": "",
                      "action_creator_social_network_id": 0,
                      "action_creator_social_network_name": "",
                      "action_creator_social_network_username": ""
                  },
                  {
                      "action_code_name": "",
                      "action_quantity": 0,
                      "action_cost_price": 0,
                      "action_cost_currency": "",
                      "action_bundle_price": 0,
                      "action_accepted_price": 0,
                      "action_draft_content_status": "",
                      "action_final_content_status": "",
                      "action_upload_draft_content_date": "",
                      "action_upload_final_content_date": "",
                      "action_payment_condition": "",
                      "action_creator_id": 0,
                      "action_creator_name": "",
                      "action_creator_avatar": "",
                      "action_creator_email": "",
                      "action_creator_social_network_id": 0,
                      "action_creator_social_network_name": "",
                      "action_creator_social_network_username": ""
                  }
              ]
          }
      ],
      "page_size": 1,
      "status": "SUCCESS",
      "total_items": 1,
      "total_pages": 1
  }
  ```

- [x] Advanced queries

  We want to search campaigns by campaign_creator_social_network_actions.codename=INSTAGRAM_POST_PHOTO

- [x] New routes

  We want to have an POST endpoint to save campaign_creator_social_network_actions

- [x] Optimization

  We want to create 2 libraries core-models-private-library and core-utils-private-library and install them in the go.mod file
