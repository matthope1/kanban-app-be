package db

// func GetBoards(db *gorm.DB, userEmail string) []types.Board {
// 	var boards []types.Board
// 	db.Raw("SELECT * from board WHERE user_email = ?", userEmail).Scan(&boards)
// 	return boards
// }

// func GetColumns(db *gorm.DB, boardId int) []types.Column {
// 	var columns []types.Column
// 	db.Raw(`SELECT * from "column" c WHERE board_id = ?`, boardId).Scan(&columns)
// 	return columns
// }

// func GetTasks(db *gorm.DB, columnId int) []types.Task {
// 	var tasks []types.Task
// 	db.Raw("SELECT * from task WHERE column_id = ?", columnId).Scan(&tasks)
// 	return tasks
// }

// func GetSubTasks(db *gorm.DB, taskId int) []types.Subtask {
// 	var subtasks []types.Subtask
// 	db.Raw("SELECT * from subtask WHERE task_id = ?", taskId).Scan(&subtasks)
// 	return subtasks
// }

// // TODO: update functions
// func UpdateBoard(db *gorm.DB, board types.Board) {
// 	// On the front end, if the user wants to update the column names, we need to update the column names in the column table
// 	if err := db.Exec("UPDATE board set title = ?, status = ? WHERE id = ?", board.Title, board.Status, board.ID).Error; err != nil {
// 		fmt.Println("error updating board:", err)
// 	}
// }

// // TODO: add functions
// func AddBoard(db *gorm.DB, board types.Board, userEmail string) (int, error) {
// 	// INSERT INTO table_name (column1, column2, column3, ...)
// 	// VALUES (value1, value2, value3, ...);
// 	// sql to insert board into board table
// 	fmt.Println("adding board to db:", board.Title, board.Status, userEmail)
// 	var ID int
// 	// TODO: we need to be careful about the default int being 0
// 	if err := db.Raw("INSERT INTO board (title, status, user_email) VALUES (?, ?, ?) RETURNING id",
// 		board.Title, board.Status, board.UserEmail).Scan(&ID).Error; err != nil {
// 		fmt.Println("error adding board to db:", err)
// 		return -1, err
// 	}
// 	fmt.Println("returned id....", ID)
// 	// TODO: return the id of the board just created and use it to add the columns
// 	return ID, nil
// }

// func AddColumn(db *gorm.DB, column types.Column, boardId int) {
// 	fmt.Println("adding column to db:", column.Title, boardId)

// 	if err := db.Exec(`INSERT INTO "column" (title, board_id) VALUES (?, ?)`,
// 		column.Title, boardId).Error; err != nil {
// 		fmt.Println("error adding column to db:", err)
// 	}
// }

// func UpdateColumn(db *gorm.DB, column types.Column) {
// 	fmt.Println("updating column in db:", column.Title)

// 	if err := db.Exec(`UPDATE "column" set title = ? WHERE id = ?`,
// 		column.Title, column.ID).Error; err != nil {
// 		fmt.Println("error updating column in db:", err)
// 	}
// }

// func AddTask(db *gorm.DB, task types.Task) {
// 	fmt.Println("adding task to db:", task.Desc, task.ColumnId, task.Desc, task.Status, task.Title)

// 	if err := db.Exec(`INSERT INTO task (status, "desc", column_id, title) VALUES (?, ?, ?, ?)`,
// 		task.Status, task.Desc, task.ColumnId, task.Title).Error; err != nil {
// 		fmt.Println("error adding task to db:", err)
// 	}
// }

// func UpdateTask(db *gorm.DB, task types.Task) {
// 	fmt.Println("updating task in db:", task.Desc, task.Status)

// 	if err := db.Exec(`UPDATE task set "desc" = ?, status = ?, title = ? WHERE id = ?`,
// 		task.Desc, task.Status, task.Title, task.ID).Error; err != nil {
// 		fmt.Println("error updating task")
// 	}
// }

// func AddSubtask(db *gorm.DB, subtask types.Subtask) {
// 	fmt.Println("adding subtask to db:", subtask.TaskId, subtask.Desc, subtask.IsComplete)

// 	if err := db.Exec(`INSERT INTO subtask (task_id, "desc", is_complete) VALUES (?, ?, ?)`,
// 		subtask.TaskId, subtask.Desc, subtask.IsComplete).Error; err != nil {
// 		fmt.Println("error adding subtask to db:", err)
// 	}
// }

// func UpdateSubtask(db *gorm.DB, subtask types.Subtask) {
// 	fmt.Println("updating subtask in db:", subtask.Desc, subtask.IsComplete)

// 	if err := db.Exec(`UPDATE subtask set "desc" = ?, is_complete = ? WHERE id = ?`,
// 		subtask.Desc, subtask.IsComplete, subtask.ID).Error; err != nil {
// 		fmt.Println("error updating subtask")
// 	}
// }

// func GetBoardOwnerById(db *gorm.DB, boardId int) string {
// 	var userEmail string
// 	db.Raw("SELECT user_email from board WHERE id = ?", boardId).Scan(&userEmail)
// 	return userEmail
// }
