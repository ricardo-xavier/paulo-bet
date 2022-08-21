package xavier.ricardo.paulobet.tasks;

import android.app.ProgressDialog;
import android.os.AsyncTask;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Call;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.RequestBody;
import okhttp3.Response;
import xavier.ricardo.paulobet.BetActivity;
import xavier.ricardo.paulobet.Config;
import xavier.ricardo.paulobet.model.BetRequest;

public class BetTask extends AsyncTask<String, Void, Response> {
    private ProgressDialog progress;
    private BetActivity context;
    private String leagueId;
    private String userId;
    private String matchId;
    private int home;
    private int visitors;
    private String token;

    private static final MediaType JSON = MediaType.parse("application/json; charset=utf-8");

    public BetTask(BetActivity context, String leagueId, String userId, String matchId, int home, int visitors, String token) {
        this.context = context;
        this.leagueId = leagueId;
        this.userId = userId;
        this.matchId = matchId;
        this.home = home;
        this.visitors = visitors;
        this.token = token;
    }

    @Override
    protected void onPreExecute() {
        progress = new ProgressDialog(context);
        progress.setMessage("Aguarde...");
        progress.show();
        super.onPreExecute();
    }

    @Override
    protected void onPostExecute(Response response) {
        super.onPostExecute(response);
        progress.dismiss();
        try {
            context.onTaskResponse(response);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    @Override
    protected void onProgressUpdate(Void... values) {
        super.onProgressUpdate(values);
    }

    @Override
    protected Response doInBackground(String... params) {
        try {
            Gson gson = new Gson();
            BetRequest betRequest = new BetRequest();
            betRequest.setMatchId(matchId);
            betRequest.setHome(home);
            betRequest.setVisitors(visitors);
            betRequest.setToken(token);
            String json = gson.toJson(betRequest, BetRequest.class);
            RequestBody body = RequestBody.create(JSON, json);
            Request request = new Request.Builder().url(Config.URL_GET_BETS + leagueId + "/" + userId).post(body).build();
            OkHttpClient client = new OkHttpClient.Builder().build();
            Call call = client.newCall(request);
            return call.execute();
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }
}
