package main

import "uni-schedule-backend/internal/app"

func main() {
	/*
				ENDPOINTS:
					SCHEDULE:
						- [GET] /schedule/id/:id
						- [GET] /schedule/slug/:slug
						- [PATCH] /schedule/:id
					LECTURERS:
						- [GET] /lecturer/all/:schedule_id
						- [PATCH] /lecturer/:id
						- [POST] /lecturer/:schedule_id
					PAIRS:
						- [POST] /pair
						- [PATCH] /pair/:id
						- [DELETE] /pair/:id
		  				- [POST] /pair/item
						- [POST] /pair/item


	*/

	appInstance := app.New()
	appInstance.Init()
}
