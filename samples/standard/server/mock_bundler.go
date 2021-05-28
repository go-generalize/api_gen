// +build mock

// Code generated by api_gen. DO NOT EDIT.
// generated version: devel
package controller

import (
	"net/http"

	types_api_0b9a5e6c "github.com/go-generalize/api_gen/samples/standard/api"
	types_service_9e7a8e17 "github.com/go-generalize/api_gen/samples/standard/api/service"
	types_static_page_de0aba5f "github.com/go-generalize/api_gen/samples/standard/api/service/static_page"
	types_user_256eedda "github.com/go-generalize/api_gen/samples/standard/api/service/user"
	types_user2_058d5a8a "github.com/go-generalize/api_gen/samples/standard/api/service/user2"
	types__userID_c13e838d "github.com/go-generalize/api_gen/samples/standard/api/service/user2/_userID"
	types__JobID_3469130f "github.com/go-generalize/api_gen/samples/standard/api/service/user2/_userID/_JobID"
	ctrl_controller_12c43f8d "github.com/go-generalize/api_gen/samples/standard/server/mock/controller"
	ctrl_service_61e5de06 "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service"
	ctrl_static_page_87ad62fe "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service/static_page"
	ctrl_user_e4c8f2cf "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service/user"
	ctrl_user2_2348a599 "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service/user2"
	ctrl__userID_145a08ca "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service/user2/_userID"
	ctrl__JobID_59b29e1f "github.com/go-generalize/api_gen/samples/standard/server/mock/controller/service/user2/_userID/_JobID"
	apierror "github.com/go-generalize/api_gen/samples/standard/server/pkg/apierror"
	props "github.com/go-generalize/api_gen/samples/standard/server/props"
	echo "github.com/labstack/echo/v4"
	xerrors "golang.org/x/xerrors"
)

func addRoutes(e *echo.Echo, p *props.ControllerProps) {
	add := func(method, path string, handler func(c echo.Context) (interface{}, error)) {
		e.Add(method, path, func(c echo.Context) error {
			var werr *apierror.APIError

			res, err := handler(c)

			if err != nil {
				if xerrors.As(err, &werr) {
					c.Logger().Errorf("%+v", werr)
					return c.JSON(werr.Status, werr.Body)
				}
				return err
			}
			if res == nil {
				return nil
			}

			return c.JSON(http.StatusOK, res)
		})
	}

	{
		ctrl := ctrl_controller_12c43f8d.NewGetController(p)

		add("GET", "/", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_api_0b9a5e6c.GetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.Get(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("Get returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_controller_12c43f8d.NewPostCreateTableController(p)

		add("POST", "/create_table", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_api_0b9a5e6c.PostCreateTableRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/create_table): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostCreateTable(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostCreateTable returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_controller_12c43f8d.NewPostCreateUserController(p)

		add("POST", "/create_user", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_api_0b9a5e6c.PostCreateUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/create_user): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostCreateUser(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostCreateUser returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_service_61e5de06.NewGetArticleController(p)

		add("GET", "/service/article", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_service_9e7a8e17.GetArticleRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/article): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.GetArticle(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("GetArticle returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_static_page_87ad62fe.NewGetStaticPageController(p)

		add("GET", "/service/static_page/static_page", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_static_page_de0aba5f.GetStaticPageRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/static_page/static_page): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.GetStaticPage(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("GetStaticPage returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_e4c8f2cf.NewGetController(p)

		add("GET", "/service/user/", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user_256eedda.GetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.Get(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("Get returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_e4c8f2cf.NewPostUpdateUserNameController(p)

		add("POST", "/service/user/update_user_name", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user_256eedda.PostUpdateUserNameRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/update_user_name): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostUpdateUserName(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostUpdateUserName returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user_e4c8f2cf.NewPostUpdateUserPasswordController(p)

		add("POST", "/service/user/update_user_password", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user_256eedda.PostUpdateUserPasswordRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user/update_user_password): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostUpdateUserPassword(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostUpdateUserPassword returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_2348a599.NewGetUserController(p)

		add("GET", "/service/user2/:user_id", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user2_058d5a8a.GetUserRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:user_id): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.GetUser(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("GetUser returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_2348a599.NewPostUpdateUserNameController(p)

		add("POST", "/service/user2/update_user_name", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user2_058d5a8a.PostUpdateUserNameRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/update_user_name): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostUpdateUserName(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostUpdateUserName returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl_user2_2348a599.NewPostUpdateUserPasswordController(p)

		add("POST", "/service/user2/update_user_password", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types_user2_058d5a8a.PostUpdateUserPasswordRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/update_user_password): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PostUpdateUserPassword(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PostUpdateUserPassword returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl__userID_145a08ca.NewGetUserJobGetController(p)

		add("GET", "/service/user2/:userID/user_job_get", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types__userID_c13e838d.GetUserJobGetRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:userID/user_job_get): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.GetUserJobGet(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("GetUserJobGet returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

	{
		ctrl := ctrl__JobID_59b29e1f.NewPutJobController(p)

		add("PUT", "/service/user2/:userID/:JobID/job", func(c echo.Context) (interface{}, error) {
			var werr *apierror.APIError

			req := new(types__JobID_3469130f.PutJobRequest)
			if err := c.Bind(req); err != nil {
				c.Logger().Errorf("failed to bind a request for (/service/user2/:userID/:JobID/job): %+v", err)
				return nil, c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":    http.StatusBadRequest,
					"message": "invalid request.",
				})
			}
			if err := c.Validate(req); err != nil && err != echo.ErrValidatorNotRegistered {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, err
			}

			res, err := ctrl.PutJob(c, req)

			if err != nil {
				if xerrors.As(err, &werr) {
					return nil, c.JSON(werr.Status, werr.Body)
				}
				return nil, xerrors.Errorf("PutJob returned an error: %w", err)
			}

			if res == nil {
				return nil, nil
			}

			return res, nil
		})
	}

}
