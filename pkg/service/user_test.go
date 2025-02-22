package service

import (
	"context"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/rand"

	"github.com/MicroOps-cn/fuck-web/pkg/service/models"
)

func testUserService(ctx context.Context, t *testing.T, svc Service) {
	var userId string
	oriUser := models.User{
		Username:    "lion",
		Email:       "lion@fuck_web.local",
		PhoneNumber: "+0112345678",
		FullName:    "Lion",
		Avatar:      "xxxxxxxxxxx",
		Status:      models.UserMeta_user_inactive,
	}

	if !t.Run("Test Create User", func(t *testing.T) {
		cUser := oriUser
		count, users, err := svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 1024)
		require.NoError(t, err)
		require.Len(t, users, 0)
		require.Equal(t, count, int64(0))
		t.Run("Test Create Null User", func(t *testing.T) {
			err = svc.CreateUser(ctx, &models.User{})
			require.Error(t, err)
		})
		for i := 0; i < 5; i++ {
			err = svc.CreateUser(ctx, &models.User{
				Username:    rand.String(5),
				Email:       rand.String(7),
				PhoneNumber: rand.String(9),
				FullName:    rand.String(5),
				Avatar:      rand.String(20),
				Status:      models.UserMeta_UserStatus(rand.Intn(4)),
			})
			require.NoError(t, err)
		}

		err = svc.CreateUser(ctx, &cUser)
		require.NoError(t, err)
		require.NotEmpty(t, cUser.Id)
		_, err = uuid.FromString(cUser.Id)
		require.NoError(t, err)
		userId = cUser.Id
		user, err := svc.GetUserInfo(ctx, cUser.Id, "")
		require.NoError(t, err)
		require.True(t, time.Since(cUser.CreateTime) < time.Second*3 && time.Since(user.CreateTime) > -time.Second)
		require.Equal(t, user.Username, "lion")
		require.Equal(t, user.FullName, "Lion")
		require.Equal(t, user.Email, "lion@fuck-web.local")
		require.Equal(t, user.PhoneNumber, "+0112345678")
		require.Equal(t, user.Avatar, "xxxxxxxxxxx")
		require.Equal(t, user.Status, models.UserMeta_user_inactive)

		t.Run("Test Create Duplicate User", func(t *testing.T) {
			cUser = oriUser
			err = svc.CreateUser(ctx, &cUser)
			require.Error(t, err)
		})
		for i := 0; i < 5; i++ {
			err = svc.CreateUser(ctx, &models.User{
				Username:    rand.String(5),
				Email:       rand.String(7),
				PhoneNumber: rand.String(9),
				FullName:    rand.String(5),
				Avatar:      rand.String(20),
				Status:      models.UserMeta_UserStatus(rand.Intn(4)),
			})
			require.NoError(t, err)
		}

		count, users, err = svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)
		require.Len(t, users, 11)
		require.Equal(t, count, int64(11))

		for _, u := range users {
			if u.Id == userId {
				_, err = uuid.FromString(u.Id)
				require.NoError(t, err)
				require.Truef(t, time.Since(u.CreateTime) <= time.Minute && time.Since(u.CreateTime) >= -time.Second, "now=%s, createTime=%s,sub=%s", time.Now(), u.CreateTime, time.Since(u.CreateTime).String())
				require.Equal(t, u.Username, "lion")
				require.Equal(t, u.FullName, "Lion")
				require.Equal(t, u.Email, "lion@fuck-web.local")
				require.Equal(t, u.PhoneNumber, "+0112345678")
				require.Equal(t, u.Avatar, "xxxxxxxxxxx")
				require.Equal(t, u.Status, models.UserMeta_user_inactive)
			}
		}
	}) {
		return
	}

	t.Run("Test Get Users", func(t *testing.T) {
		count, users, err := svc.GetUsers(ctx, "Asdooa299shdoiasgd8269bw3i7y9fdsahigf", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)
		require.Len(t, users, 0)
		require.Equal(t, count, int64(0))

		_, users, err = svc.GetUsers(ctx, "", models.UserMeta_user_inactive, 1, 20)
		require.NoError(t, err)
		for _, user := range users {
			require.Equal(t, user.Status, models.UserMeta_user_inactive)
		}

		_, users, err = svc.GetUsers(ctx, "", models.UserMeta_normal, 1, 20)
		require.NoError(t, err)
		for _, user := range users {
			require.Equal(t, user.Status, models.UserMeta_normal)
		}

		_, users, err = svc.GetUsers(ctx, "", oriUser.Status, 1, 20)
		require.NoError(t, err)
		found := false
		for _, user := range users {
			require.Equal(t, user.Status, oriUser.Status)
			if user.Id == userId {
				found = true
			}
		}
		require.Equal(t, found, true)
	})

	if !t.Run("Test Update User", func(t *testing.T) {
		oriUser1 := &models.User{
			Model:       models.Model{Id: userId},
			Username:    "lion_u",
			Email:       "lion_u@fuck-web.local",
			PhoneNumber: "+01123456789",
			FullName:    "Lion_u",
			Avatar:      "xxxxxxxxxxx_u",
			Status:      models.UserMeta_normal,
		}
		err := svc.UpdateUser(ctx, oriUser1)
		require.NoError(t, err)
		user, err := svc.GetUserInfo(ctx, oriUser1.Id, "")
		require.NoError(t, err)
		_, err = uuid.FromString(user.Id)
		require.NoError(t, err)
		require.Truef(t, time.Since(user.CreateTime) <= time.Minute && time.Since(user.CreateTime) >= -time.Second, "now=%s, createTime=%s,sub=%s", time.Now(), user.CreateTime, time.Since(user.CreateTime).String())
		require.Equal(t, user.Username, "lion")
		require.Equal(t, user.FullName, "Lion_u")
		require.Equal(t, user.Email, "lion_u@fuck-web.local")
		require.Equal(t, user.PhoneNumber, "+01123456789")
		require.Equal(t, user.Avatar, "xxxxxxxxxxx_u")
		require.Equal(t, user.Status, models.UserMeta_normal)
		count, users, err := svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)
		require.Len(t, users, 11)
		require.Equal(t, count, int64(11))

		for _, u := range users {
			if u.Id == userId {
				_, err = uuid.FromString(u.Id)
				require.NoError(t, err)
				require.True(t, time.Since(u.CreateTime) < time.Second*3 && time.Since(u.CreateTime) > -time.Second)
				require.Equal(t, u.Username, "lion")
				require.Equal(t, u.FullName, "Lion_u")
				require.Equal(t, u.Email, "lion_u@fuck-web.local")
				require.Equal(t, u.PhoneNumber, "+01123456789")
				require.Equal(t, u.Avatar, "xxxxxxxxxxx_u")
				require.Equal(t, u.Status, models.UserMeta_normal)
			}
		}
	}) {
		return
	}

	t.Run("Test Update some fields of users", func(t *testing.T) {
		oriUser1 := &models.User{
			Model:       models.Model{Id: userId},
			Username:    "lion_u2",
			Email:       "lion_u2@fuck-web.local",
			PhoneNumber: "+011234567890",
			FullName:    "Lion_u2",
			Avatar:      "xxxxxxxxxxx_u2",
			Status:      models.UserMeta_user_inactive,
		}
		err := svc.UpdateUser(ctx, oriUser1, "email", "avatar")
		require.NoError(t, err)
		_, err = uuid.FromString(oriUser1.Id)
		require.NoError(t, err)

		user, err := svc.GetUserInfo(ctx, oriUser1.Id, "")
		require.NoError(t, err)
		require.True(t, time.Since(user.CreateTime) < time.Second*3 && time.Since(user.CreateTime) > -time.Second)
		require.Equal(t, user.Username, "lion")
		require.Equal(t, user.FullName, "Lion_u")
		require.Equal(t, user.Email, "lion_u2@fuck-web.local")
		require.Equal(t, user.PhoneNumber, "+01123456789")
		require.Equal(t, user.Avatar, "xxxxxxxxxxx_u2")
		require.Equal(t, user.Status, models.UserMeta_normal)
		count, users, err := svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)
		require.Len(t, users, 11)
		require.Equal(t, count, int64(11))

		for _, u := range users {
			if u.Id == userId {
				_, err = uuid.FromString(u.Id)
				require.NoError(t, err)
				require.True(t, time.Since(u.CreateTime) < time.Second*3 && time.Since(u.CreateTime) > -time.Second)
				require.Equal(t, u.Username, "lion")
				require.Equal(t, u.FullName, "Lion_u")
				require.Equal(t, u.Email, "lion_u2@fuck-web.local")
				require.Equal(t, u.PhoneNumber, "+01123456789")
				require.Equal(t, u.Avatar, "xxxxxxxxxxx_u2")
				require.Equal(t, u.Status, models.UserMeta_normal)
			}
		}
	})
	t.Run("Test Patch User", func(t *testing.T) {
		err := svc.PatchUser(ctx, map[string]interface{}{"id": userId, "status": models.UserMeta_disabled})
		require.NoError(t, err)

		user, err := svc.GetUserInfo(ctx, userId, "")
		require.Equal(t, user.Status, models.UserMeta_disabled)
		require.NoError(t, err)
		_, users, err := svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)

		require.Len(t, users, 11)
		for _, u := range users {
			if u.Id == userId {
				require.Equal(t, u.Status, models.UserMeta_disabled)
			}
		}
	})

	t.Run("Test Delete User", func(t *testing.T) {
		err := svc.DeleteUser(ctx, userId)
		require.NoError(t, err)
		count, users, err := svc.GetUsers(ctx, "", models.UserMetaStatusAll, 1, 20)
		require.NoError(t, err)
		require.Len(t, users, 10)
		require.Equal(t, count, int64(10))
	})
}
