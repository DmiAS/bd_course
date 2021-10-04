package handler

//func (h *Handler) createThread(ctx echo.Context) error{
//	req := &info{}
//	if err := req.bind(ctx); err != nil {
//		return ctx.String(http.StatusBadRequest, err.Error())
//	}
//
//
//	if err := h.threads.Create(thread); err != nil {
//		ctx.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	resp := converters2.ConvertCreateProjectOutput(id, thread.Name)
//
//	ctx.JSON(http.StatusOK, resp)
//
//}

//func (h *Handler) getThreads(ctx *gin.Context) {
//	id, err := extractID(ctx)
//	if err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	threads, err := h.threads.Get(id)
//	if err != nil {
//		ctx.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	resp := converters2.ConvertGetThreadsOutput(threads)
//
//	ctx.JSON(http.StatusOK, resp)
//}
//
//func (h *Handler) updateThread(ctx *gin.Context) {
//	id, err := extractID(ctx)
//	if err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	req := new(ds.UpdateThreadInput)
//	if err := ctx.BindJSON(req); err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	thread := converters2.ConvertUpdateThreadInput(req.Name, id)
//
//	if err := h.threads.Update(thread); err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//}
//
//func (h *Handler) deleteThread(ctx *gin.Context) {
//	id, err := extractID(ctx)
//	if err != nil {
//		ctx.String(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	if err := h.threads.Delete(id); err != nil {
//		ctx.String(http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	ctx.Status(http.StatusOK)
//}
